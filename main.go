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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

// MongoDB setup
var client *mongo.Client
var userCollection *mongo.Collection

const dbName = "testdb"
const userCollectionName = "users"
const albumCollectionName = "albums"
const songCollectionName = "songs"
const secretKey = "The_Dark_Side_Of_The_Moon"

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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

// JWT Claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {

	host := "127.0.0.1:8080"

	fmt.Println("Starting server on " + host)

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, _ = mongo.Connect(context.TODO(), clientOptions)

	// Start server
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}

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



// Register handler
func registerHandler(w http.ResponseWriter, r *http.Request) {

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
