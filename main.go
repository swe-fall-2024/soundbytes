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
	PostID        int         `bson:"_id,omitempty" json:"post_id"`
	User          string      `bson:"user" json:"user"`
	Profile_Image string      `bson:"profile_img" json:"profile_img"`
	Type          string      `bson:"type" json:"type"`
	Title         string      `bson:"title" json:"title"`
	Content       PostContent `bson:"content" json:"content"`
	LikeCount     int         `bson:"like_count" json:"like_count"`
}

type PostContent struct {
	SongTitle     string `bson:"song_title,omitempty" json:"song_title,omitempty"`
	SongURL       string `bson:"song_url,omitempty" json:"song_url,omitempty"`
	SongEmbed     string `bson:"song_embed,omitempty" json:"song_embed,omitempty"`
	AlbumTitle    string `bson:"album_title,omitempty" json:"album_title,omitempty"`
	Review        string `bson:"review,omitempty" json:"review,omitempty"`
	PlaylistTitle string `bson:"playlist_title,omitempty" json:"playlist_title,omitempty"`
	PlaylistURL   string `bson:"playlist_url,omitempty" json:"playlist_url,omitempty"`
	PlaylistEmbed string `bson:"playlist_embed,omitempty" json:"playlist_embed,omitempty"`
}

// // User struct
// type User struct {
// 	UserID    primitive.ObjectID `bson:"_id,omitempty" json:"user_id"`
// 	Username  string             `bson:"username" json:"username"`
// 	Password  string             `json:"password"`
// 	TopArtist string             `bson:"top_artist" json:"top_artist"`
// 	TopSong   string             `bson:"top_song" json:"top_song"`
// 	FavSongs  []string           `bson:"favorite_songs" json:"favorite_songs"`
// 	FavGenres []string           `bson:"favorite_genres" json:"favorite_genres"`
// 	Posts     []Post             `bson:"posts" json:"posts"`
// 	Following []string           `json:"following"` // List of usernames the user follows
// }

type User struct {
	UserID    string   `bson:"_id,omitempty" json:"userID"` // Change to match Angular `userID`
	Email     string   `bson:"email" json:"email"`
	Username  string   `bson:"username" json:"username"`
	Password  string   `json:"password"`
	TopArtist string   `bson:"top_artist" json:"topArtist"`      // Match Angular `topArtist`
	TopSong   string   `bson:"top_song" json:"topSong"`          // Match Angular `topSong`
	FavSongs  []string `bson:"favorite_songs" json:"favSongs"`   // Match Angular `favSongs`
	FavGenres []string `bson:"favorite_genres" json:"favGenres"` // Match Angular `favGenres`
	Posts     []Post   `bson:"posts" json:"posts"`
	Following []string `json:"following"`
}

// Initialize the User struct
func createUser(user *User) {
	// Set default values for empty slices
	if user.Email == "" {
		user.Email = "jo mama"
	}

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

	postCollection := client.Database(dbName).Collection("posts")

	var post Post
	// Decode the request body into the Post struct
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Check for required fields in the post (assuming title and content are required)
	if post.Title == "" {
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
	router.HandleFunc("/profile", updateUserProfile).Methods("PUT")
	router.HandleFunc("/getPosts/{username}", getPostsHandler).Methods("GET")

	// Protect this route with JWT middleware
	router.HandleFunc("/protected", jwtMiddleware(protectedHandler)).Methods("GET")

	// Search Bar Handler
	router.HandleFunc("/searchUsers", searchUsersHandler).Methods("GET")

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
	fmt.Println("Username: ", user.Username) // Log if no userID is provided
	fmt.Println("Password: ", user.Password) // Log if no userID is provided

	err := userCollection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&foundUser)
	if err != nil {
		fmt.Println("failed early: ", user.Password) // Log if no userID is provided

		http.Error(w, "Invalid credentialssssss", http.StatusUnauthorized)
		return
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		fmt.Println("fawiled ttttt: ", user.Password) // Log if no userID is provided
		fmt.Println("fawiled wewewe: ", foundUser.Password) // Log if no userID is provided

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

func updateUserProfile(w http.ResponseWriter, r *http.Request) {
	// Get the userID from query parameters
	userID := r.URL.Query().Get("userId") // Extract userId from query string

	// Ensure userID is provided
	if userID == "" {
		fmt.Println("Error: No userID provided") // Log if no userID is provided
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userCollection := client.Database(dbName).Collection(userCollectionName)

	fmt.Println("Received userID:", userID) // Log the user ID to check if it's passed correctly

	// Log the request details
	fmt.Println("Request Body:", r.Body)

	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		fmt.Println("Error decoding request body:", err) // Log the error message
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	fmt.Println("Decoded user data:", updatedUser) // Log the decoded user data

	// Prepare the update data
	updateData := bson.M{
		"$set": bson.M{
			"email":		   updatedUser.Email, 
			"username":        updatedUser.Username,
			"password":        updatedUser.Password, // Ensure this is a hashed password
			"top_artist":      updatedUser.TopArtist,
			"top_song":        updatedUser.TopSong,
			"favorite_songs":  updatedUser.FavSongs,
			"favorite_genres": updatedUser.FavGenres,
			"posts":           updatedUser.Posts,     // Ensure that posts are handled correctly (null or empty array)
			"following":       updatedUser.Following, // Ensure that following is handled correctly (null or empty array)
		},
	}

	// If any of the fields are null or empty, handle them appropriately
	if updatedUser.FavSongs == nil {
		updateData["$set"].(bson.M)["favorite_songs"] = []string{}
	}
	if updatedUser.FavGenres == nil {
		updateData["$set"].(bson.M)["favorite_genres"] = []string{}
	}
	if updatedUser.Posts == nil {
		updateData["$set"].(bson.M)["posts"] = []interface{}{}
	}
	if updatedUser.Following == nil {
		updateData["$set"].(bson.M)["following"] = []string{}
	}

	// Log the update data for debugging
	fmt.Println("Update data for MongoDB:", updateData)

	// Update the user in the database
	filter := bson.M{"username": userID} // Use the userID to find the user

	result, err := userCollection.UpdateOne(context.TODO(), filter, updateData)
	if err != nil {
		fmt.Println("Error updating profile:", err) // Log the error if update fails
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		fmt.Println("No matching user found for update") // Log if no match was found
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}

	fmt.Println("Profile updated successfully") // Log if update is successful
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
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
	collection := client.Database(dbName).Collection(userCollectionName)

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
		"userID":    user["_id"],             // map _id to userID
		"email": 	 user["email"],
		"username":  user["username"],        // map username to username
		"password":  user["password"],        // map password to password
		"topArtist": user["top_artist"],      // rename top_artist to topArtist
		"topSong":   user["top_song"],        // rename top_song to topSong
		"favSongs":  user["favorite_songs"],  // rename favorite_songs to favSongs
		"favGenres": user["favorite_genres"], // rename favorite_genres to favGenres
		"posts":     user["posts"],           // posts can stay the same
		"following": user["following"],       // following can stay the same
	}

	// Encode the response and send it to the client
	json.NewEncoder(w).Encode(response)
}

func getPostsHandler(w http.ResponseWriter, r *http.Request) {
	// Use a different collection for posts
	postCollection := client.Database(dbName).Collection("posts")

	// Get the username from the request URL
	vars := mux.Vars(r)
	user := vars["username"]

	// Ensure the username is provided
	if user == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Find posts by the username
	cursor, err := postCollection.Find(context.TODO(), bson.M{"user": user})

	if err != nil {
		http.Error(w, "Posts not found", http.StatusNotFound)
		return
	}
	defer cursor.Close(context.TODO())

	var posts []Post

	for cursor.Next(context.TODO()) {
		var post Post
		if err := cursor.Decode(&post); err != nil {
			http.Error(w, "Error decoding post", http.StatusInternalServerError)
			return
		}
		posts = append(posts, post)
	}

	log.Println(fmt.Sprintf("Posts %v", posts))

	// Return the user's posts
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
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

func getProfileHandlerForTesting(w http.ResponseWriter, r *http.Request, client *mongo.Client, userCollection *mongo.Collection) {

	userCollection = client.Database(dbName).Collection(userCollectionName)

	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	// Find user in MongoDB
	var user struct {
		Username         string   `json:"username"`
		Name             string   `json:"name"`
		TopArtist        string   `json:"top_artist"`
		TopSong          string   `json:"top_song"`
		TopGenre         string   `json:"top_genre"`
		Top3Genres       []string `json:"top_3_genres"`
		AllTimeFavSong   string   `json:"all_time_fav_song"`
		AllTimeFavArtist string   `json:"all_time_fav_artist"`
	}

	err := userCollection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Return the user's profile
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func registerSongHandlerForTesting(w http.ResponseWriter, r *http.Request, client *mongo.Client, userCollection *mongo.Collection) {

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

	client1 := spotify.Authenticator{}.NewClient(accessToken)

	var song1 Song

	json.NewDecoder(r.Body).Decode(&song1)

	log.Println(song1.Id)

	songID := spotify.ID(song1.Id)

	song, err := client1.GetTrack(songID)

	if err != nil {
		log.Fatalf("error retrieve playlist data: %v", err)
	}

	song1.Artist = song.Artists[0].Name
	song1.Name = song.Name
	song1.Popularity = string(rune(song.Popularity))

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

func registerAlbumHandlerForTesting(w http.ResponseWriter, r *http.Request, client *mongo.Client, userCollection *mongo.Collection) {

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

	client1 := spotify.Authenticator{}.NewClient(accessToken)

	var album1 Album

	json.NewDecoder(r.Body).Decode(&album1)

	albumID := spotify.ID(album1.Id)

	album, err := client1.GetAlbum(albumID)

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

func registerHandlerForTesting(w http.ResponseWriter, r *http.Request, client *mongo.Client, userCollection *mongo.Collection) {

	userCollection = client.Database(dbName).Collection(userCollectionName)

	var user User

	json.NewDecoder(r.Body).Decode(&user)

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
func loginHandlerForTesting(w http.ResponseWriter, r *http.Request, client *mongo.Client, userCollection *mongo.Collection) {

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

func followUserHandlerForTesting(w http.ResponseWriter, r *http.Request, client *mongo.Client, userCollection *mongo.Collection) {

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
func setUpProfileForTesting(w http.ResponseWriter, r *http.Request, client *mongo.Client, userCollection *mongo.Collection) {

	userCollection = client.Database(dbName).Collection(userCollectionName)

	// Decode request body
	var request struct {
		Username         string   `json:"username"`
		Name             string   `json:"name"`
		TopArtist        string   `json:"top_artist"`
		TopSong          string   `json:"top_song"`
		TopGenre         string   `json:"top_genre"`
		Top3Genres       []string `json:"top_3_genres"`
		AllTimeFavSong   string   `json:"all_time_fav_song"`
		AllTimeFavArtist string   `json:"all_time_fav_artist"`
	}

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {

		http.Error(w, "Invalid request", http.StatusBadRequest)
		return

	}

	filter := bson.M{"username": request.Username}

	// Define update operation (set new values)
	update := bson.M{
		"$set": bson.M{
			"name":                request.Name,
			"top_artist":          request.TopArtist,
			"top_song":            request.TopSong,
			"top_genre":           request.TopGenre,
			"top_3_genres":        request.Top3Genres,
			"all_time_fav_song":   request.AllTimeFavSong,
			"all_time_fav_artist": request.AllTimeFavArtist,
		},
	}

	// Perform the update
	_, err = userCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		http.Error(w, "Error updating profile", http.StatusInternalServerError)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})

}

func unfollowUserHandlerForTesting(w http.ResponseWriter, r *http.Request, client *mongo.Client, userCollection *mongo.Collection) {

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

	update := bson.M{"$pull": bson.M{"following": request.Followee}} // Ensure no duplicates

	_, err = userCollection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		http.Error(w, "Error updating follow list", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "User unfollowed successfully"})

}

func searchUsersHandler(w http.ResponseWriter, r *http.Request) {
    userCollection = client.Database(dbName).Collection(userCollectionName)

    // Get the search query from the URL parameters
    query := r.URL.Query().Get("q")
    if query == "" {
        http.Error(w, "Query parameter 'q' is required", http.StatusBadRequest)
        return
    }

    // Use a regex search to find users whose usernames start with the query string
    filter := bson.M{"username": bson.M{"$regex": "^" + query, "$options": "i"}} // Case-insensitive search
    options := options.Find().SetLimit(3) // Limit results to 3

    cursor, err := userCollection.Find(context.TODO(), filter, options)
    if err != nil {
        http.Error(w, "Error fetching users", http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.TODO())

    // Collect the results
    var users []User
    for cursor.Next(context.TODO()) {
        var user User
        if err := cursor.Decode(&user); err != nil {
            http.Error(w, "Error decoding user", http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    // Extract only the usernames
    var usernames []string
    for _, user := range users {
        usernames = append(usernames, user.Username)
    }

    // Respond with the JSON list of usernames
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(usernames)
}
