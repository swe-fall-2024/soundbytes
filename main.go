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
)

// MongoDB setup
var client *mongo.Client
var userCollection *mongo.Collection

const dbName = "testdb"
const collectionName = "users"
const secretKey = "The_Dark_Side_Of_The_Moon"

// User struct
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
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

	userCollection = client.Database(dbName).Collection(collectionName)

	// Start server
	if err := http.ListenAndServe(host, httpHandler()); err != nil {
		log.Fatalf("Failed to listen on %s: %v", host, err)
	}
}

// httpHandler creates the backend HTTP router
func httpHandler() http.Handler {
	router := mux.NewRouter()

	// Authentication routes
	router.HandleFunc("/register", registerHandler).Methods("POST")
	router.HandleFunc("/login", loginHandler).Methods("POST")

	// Protect this route with JWT middleware (Example)
	router.HandleFunc("/protected", jwtMiddleware(protectedHandler)).Methods("GET")

	// Serve Angular app
	router.PathPrefix("/").Handler(AngularHandler).Methods("GET")

	return handlers.LoggingHandler(os.Stdout,
		handlers.CORS(
			handlers.AllowCredentials(),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"http://localhost:4200"}),
		)(router))
}

// Register handler
func registerHandler(w http.ResponseWriter, r *http.Request) {
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
	origin, _ := url.Parse("http://localhost:4200")
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
