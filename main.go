package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// MongoDB setup
var client *mongo.Client
var userCollection *mongo.Collection
var friendCollection *mongo.Collection
var postCollection *mongo.Collection

const dbName = "testdb"
const userCollectionName = "users"
const albumCollectionName = "albums"
const songCollectionName = "songs"
const secretKey = "The_Dark_Side_Of_The_Moon"

type Album struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Artist      string `json:"artist"`
	Genre       string `json:"genre"`
	ReleaseDate string `json:"release_date"`
}

type Song struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Artist     string `json:"artist"`
	Popularity string `json:"popularity"`
}

// Friend struct
type Friend struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username       string             `bson:"username" json:"username"`
	FriendUsername string             `bson:"friend_username" json:"friend_username"`
}

// Post struct
type Post struct {
	PostID    primitive.ObjectID `bson:"_id,omitempty" json:"post_id"`
	PostType  string             `bson:"post_type" json:"post_type"`
	Link      string             `bson:"link" json:"link"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	LikeCount int                `bson:"like_count" json:"like_count"`
}

// User struct
type User struct {
	UserID    primitive.ObjectID `bson:"_id,omitempty" json:"user_id"`
	Username  string             `bson:"username" json:"username"`
	Password  string             `json:"password"`
	TopArtist string             `bson:"top_artist" json:"top_artist"`
	TopSong   string             `bson:"top_song" json:"top_song"`
	FavSongs  []string           `bson:"favorite_songs" json:"favorite_songs"`
	FavGenres []string           `bson:"favorite_genres" json:"favorite_genres"`
	Posts     []Post             `bson:"posts" json:"posts"`
	Following []string           `json:"following"` // List of usernames the user follows
}


// Initialize the User struct
func createUser(user *User) {
	// Set default values for empty slices
	if user.FavSongs == nil {
		user.FavSongs = []string{}
	}
	if user.FavGenres == nil {
		user.FavGenres = []string{}
	}
	if user.Posts == nil {
		user.Posts = []Post{}
	}
	if user.Following == nil {
		user.Following = []string{}
	}
}

// JWT Claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {

	host := "127.0.0.1:4201"

	fmt.Println("Starting server on " + host)

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, _ = mongo.Connect(context.TODO(), clientOptions)

	// Start server
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}

}

// Function to create a new friend entry
func addFriend(w http.ResponseWriter, r *http.Request) {
	friendCollection := client.Database(dbName).Collection("friends")

	var friend Friend
	if err := json.NewDecoder(r.Body).Decode(&friend); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check if the user exists
	var user User // Make sure to define the User struct accordingly
	err := userCollection.FindOne(context.TODO(), bson.M{"username": friend.Username}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Check if the friend exists
	var friendUser User // Use a different variable to avoid shadowing
	err = userCollection.FindOne(context.TODO(), bson.M{"username": friend.FriendUsername}).Decode(&friendUser)
	if err != nil {
		http.Error(w, "Friend not found", http.StatusNotFound)
		return
	}

	// Check if the user is trying to add themselves
	if friend.Username == friend.FriendUsername {
		http.Error(w, "Cannot add yourself as a friend", http.StatusBadRequest)
		return
	}

	// Insert the friend entry into the database
	_, err = friendCollection.InsertOne(context.TODO(), friend)
	if err != nil {
		http.Error(w, "Failed to add friend", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Friend added successfully"})
}

// Function to create a new post
func addPost(w http.ResponseWriter, r *http.Request) {
	postCollection := client.Database(dbName).Collection("posts") // Use := to define a new variable

	var post Post
	// Decode the request body into the Post struct
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check for required fields in the post (assuming title and content are required)
	if post.Title == "" || post.Content == "" {
		http.Error(w, "Title and content are required", http.StatusBadRequest)
		return
	}

	// Insert the post into the database
	_, err := postCollection.InsertOne(context.TODO(), post)
	if err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	// Set the response header and encode the success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Post created successfully"})
}

func registerAlbumHandler(w http.ResponseWriter, r *http.Request) {

	userCollection = client.Database(dbName).Collection(albumCollectionName)

	authConfig := &clientcredentials.Config{
		ClientID:     "3bd23353135d447b80b5e0c9d70775dc",
		ClientSecret: "05454a9a1a2b48a19769e51b025798c4",
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("error retrieve access token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)

	var album1 Album

	json.NewDecoder(r.Body).Decode(&album1)

	albumID := spotify.ID(album1.Id)

	album, err := client.GetAlbum(albumID)

	if err != nil {
		log.Fatalf("error retrieve playlist data: %v", err)
	}

	album1.Artist = album.Artists[0].Name
	album1.Name = album.Name
	//album1.Genre = album.Genres[0]
	album1.ReleaseDate = album.ReleaseDate

	// Store user in MongoDB
	_, err1 := userCollection.InsertOne(context.TODO(), album1)

	if err1 != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{"message": "Album Post registered successfully"})

	log.Println("Album ID: ", album.ID)
	log.Println("Album Name: ", album.Name)
	log.Println("Artist: ", album.Artists)
	log.Println("Genre: ", album.Genres)
	log.Println("Cover Art: ", album.Images)
	log.Println("Release Date: ", album.ReleaseDate)
}

func registerSongHandler(w http.ResponseWriter, r *http.Request) {

	userCollection = client.Database(dbName).Collection(songCollectionName)

	authConfig := &clientcredentials.Config{
		ClientID:     "3bd23353135d447b80b5e0c9d70775dc",
		ClientSecret: "05454a9a1a2b48a19769e51b025798c4",
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("error retrieve access token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)

	var song1 Song

	json.NewDecoder(r.Body).Decode(&song1)

	log.Println(song1.Id)

	songID := spotify.ID(song1.Id)

	song, err := client.GetTrack(songID)

	if err != nil {
		log.Fatalf("error retrieve playlist data: %v", err)
	}

	song1.Artist = song.Artists[0].Name
	song1.Name = song.Name
	song1.Popularity = string(song.Popularity)

	// Store user in MongoDB
	_, err1 := userCollection.InsertOne(context.TODO(), song1)

	if err1 != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{"message": "Song Post registered successfully"})

	log.Println("Song ID: ", song.ID)
	log.Println("Song Name: ", song.Name)
	log.Println("Song Artist: ", song.Artists)
}

// httpHandler creates the backend HTTP router
func httpHandler() http.Handler {
	fmt.Print("inside of httpHandler in Go")

	router := mux.NewRouter()

	// ✅ Define /api/message endpoint
	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Hello from Go!"}`))
	}).Methods("GET")

	// Authentication routes
	router.HandleFunc("/register", registerHandler).Methods("POST")
	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/postAlbum", registerAlbumHandler).Methods("POST")
	router.HandleFunc("/postSong", registerSongHandler).Methods("POST")

	// Follow Friends routes
	router.HandleFunc("/follow", followUserHandler).Methods("POST")
	router.HandleFunc("/following/{username}", getFollowingHandler).Methods("GET")

	// Table Routes
	router.HandleFunc("/addFriend", addFriend).Methods("POST")
	router.HandleFunc("/addPost", addPost).Methods("POST")
	router.HandleFunc("/profile", getUserProfile).Methods("GET")

	// Protect this route with JWT middleware
	router.HandleFunc("/protected", jwtMiddleware(protectedHandler)).Methods("GET")

	// Serve Angular app
	router.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization",
				"DNT", "Keep-Alive", "User-Agent", "X-Requested-With", "If-Modified-Since",
				"Cache-Control", "Content-Range", "Range"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:4200"}),
			handlers.ExposedHeaders([]string{"DNT", "Keep-Alive", "User-Agent",
				"X-Requested-With", "If-Modified-Since", "Cache-Control",
				"Content-Type", "Content-Range", "Range", "Content-Disposition"}),
			handlers.MaxAge(86400),
		)(router))
}

// TODO: Fix this so it accounts for all of the fields in our USER TABLE

// Register handler
func registerHandler(w http.ResponseWriter, r *http.Request) {

	userCollection = client.Database(dbName).Collection(userCollectionName)

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	// Ensure that fields with slices are initialized
	createUser(&user)
	
	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	// Store user in MongoDB
	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, "Error registering user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
}

// Login handler
func loginHandler(w http.ResponseWriter, r *http.Request) {

	userCollection = client.Database(dbName).Collection(userCollectionName)

	var user User
	json.NewDecoder(r.Body).Decode(&user)

	// Fetch user from MongoDB
	var foundUser User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&foundUser)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := generateJWT(user.Username)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// Generate JWT
func generateJWT(username string) string {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secretKey))
	return tokenString
}

// JWT Middleware
func jwtMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Token is valid, proceed to the protected handler
		next.ServeHTTP(w, r)
	}
}

// Protected route example
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the protected route!"})
}

func followUserHandler(w http.ResponseWriter, r *http.Request) {
	userCollection = client.Database(dbName).Collection(userCollectionName)

	// Decode request body
	var request struct {
		Follower string `json:"follower"`
		Followee string `json:"followee"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Update the follower's document in MongoDB to add the followee to the "Following" list
	filter := bson.M{"username": request.Follower}
	update := bson.M{"$addToSet": bson.M{"following": request.Followee}} // Ensure no duplicates

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, "Error updating follow list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User followed successfully"})
}

func getFollowingHandler(w http.ResponseWriter, r *http.Request) {
	userCollection = client.Database(dbName).Collection(userCollectionName)

	// Extract username from URL parameters
	vars := mux.Vars(r)
	username := vars["username"]

	// Find user
	var user User
	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return the list of followed users
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string][]string{"following": user.Following})
}

// Handler to fetch user profile by ID
func getUserProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Assume user ID is passed as a query parameter (e.g., /profile?userId=123)
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	// Access the users collection
	collection := client.Database(dbName).Collection("users")

	// Find user by ID
	var user bson.M
	err := collection.FindOne(context.TODO(), bson.M{"username": userID}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else {
		fmt.Print("User found")
	}

	fmt.Print("what is being encoded: ", user)

	// Manually map fields to the desired format
	response := map[string]interface{}{
		"userID":    user["_id"],           // map _id to userID
		"username":  user["username"],      // map username to username
		"password":  user["password"],      // map password to password
		"topArtist": user["top_artist"],    // rename top_artist to topArtist
		"topSong":   user["top_song"],      // rename top_song to topSong
		"favSongs":  user["favorite_songs"], // rename favorite_songs to favSongs
		"favGenres": user["favorite_genres"], // rename favorite_genres to favGenres
		"posts":     user["posts"],         // posts can stay the same
		"following": user["following"],     // following can stay the same
	}

	// Encode the response and send it to the client
	json.NewEncoder(w).Encode(response)
}

// Reverse proxy for Angular app
func getOrigin() *url.URL {
	origin, err := url.Parse("http://localhost:4200")
	if err != nil {
		log.Fatalf("Failed to parse origin URL: %v", err)
	}
	return origin
}

var origin = getOrigin()
var director = func(req *http.Request) {
	req.Header.Add("X-Forwarded-Host", req.Host)
	req.Header.Add("X-Origin-Host", origin.Host)
	req.URL.Scheme = "http"
	req.URL.Host = origin.Host
}

var AngularHandler = &httputil.ReverseProxy{Director: director}
