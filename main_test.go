package main

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var testClient *mongo.Client
var testUserCollection *mongo.Collection

func setup() {
	// Setup MongoDB connection for the test
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	testClient, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

}

func teardown() {
	// Clean up test data after the tests
	testUserCollection.DeleteMany(context.TODO(), bson.M{})
}

func TestRegisterHandler(t *testing.T) {

	// Setup the test environment
	setup()

	testUserCollection = testClient.Database("testdb").Collection("users")

	// Define the test user
	testUser := User{
		Username:  "testuser234",
		Password:  "testpassword",
		Following: []string{},
	}

	// Convert testUser to JSON
	userJSON, err := json.Marshal(testUser)

	println(string(userJSON))

	if err != nil {
		t.Fatalf("Failed to marshal test user: %v", err)
	}

	// Create a new HTTP POST request with the test user data
	req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewReader(userJSON))

	println(req)

	w := httptest.NewRecorder()

	println(w)

	// Call the register handler
	registerHandlerForTesting(w, req, testClient, testUserCollection)

	// Check the response status code
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body message
	expectedMessage := `{"message": "User registered successfully"}`
	assert.JSONEq(t, expectedMessage, w.Body.String())

	// Verify that the user was actually added to the database
	var storedUser User

	err = testUserCollection.FindOne(context.TODO(), bson.M{"username": testUser.Username}).Decode(&storedUser)

	if err != nil {
		t.Fatalf("Failed to find the registered user in the database: %v", err)
	}

	// Check that the stored user has the expected username
	assert.Equal(t, testUser.Username, storedUser.Username)

	// Check that the password is hashed
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(testUser.Password))
	if err != nil {
		t.Fatalf("Password hash mismatch: %v", err)
	}
}

func TestLoginHandler(t *testing.T) {

	// Setup the test environment
	setup()

	testUserCollection = testClient.Database("testdb").Collection("users")

	// Define the test user
	testUser := User{
		Username: "testuser",
		Password: "testpassword",
	}

	// Convert testUser to JSON
	userJSON, err := json.Marshal(testUser)

	println(string(userJSON))

	if err != nil {
		t.Fatalf("Failed to marshal test user: %v", err)
	}

	// Create a new HTTP POST request with the test user data
	req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewReader(userJSON))

	println(req)

	w := httptest.NewRecorder()

	println(w)

	// Call the register handler
	loginHandlerForTesting(w, req, testClient, testUserCollection)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestFollowUserHandler(t *testing.T) {

	// Setup the test environment
	setup()

	testUserCollection = testClient.Database("testdb").Collection("users")

	// Decode request body
	var request struct {
		Follower string `json:"follower"`
		Followee string `json:"followee"`
	}

	request.Follower = "testuser"
	request.Followee = "testuser234"

	// Convert testUser to JSON
	requestJSON, err := json.Marshal(request)

	println(string(requestJSON))

	if err != nil {
		t.Fatalf("Failed to marshal test user: %v", err)
	}

	// Create a new HTTP POST request with the test user data
	req := httptest.NewRequest(http.MethodPost, "/follow", bytes.NewReader(requestJSON))

	println(req)

	w := httptest.NewRecorder()

	println(w)

	// Call the register handler
	followUserHandlerForTesting(w, req, testClient, testUserCollection)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body message
	expectedMessage := `{"message": "User followed successfully"}`
	assert.JSONEq(t, expectedMessage, w.Body.String())

}

func TestSetUpProfileHandler(t *testing.T) {

	// Setup the test environment
	setup()

	testUserCollection = testClient.Database("testdb").Collection("users")

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

	request.Username = "testuser"
	request.Name = "testname"
	request.TopArtist = "testtopartist"
	request.TopSong = "testtopsong"
	request.TopGenre = "testtopgenre"
	request.Top3Genres = []string{"testtopgenre1", "testtopgenre2", "testtopgenre3"}
	request.AllTimeFavSong = "testalltimefavsong"
	request.AllTimeFavArtist = "testalltimefavartist"

	// Convert testUser to JSON
	userJSON, err := json.Marshal(request)

	println(string(userJSON))

	if err != nil {
		t.Fatalf("Failed to marshal test user: %v", err)
	}

	// Create a new HTTP POST request with the test user data

	req := httptest.NewRequest(http.MethodPut, "/setUpProfile", bytes.NewReader(userJSON))

	println(req)

	w := httptest.NewRecorder()

	println(w)

	// Call the register handler

	setUpProfileForTesting(w, req, testClient, testUserCollection)

	// Check the response status code

	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body message

	expectedMessage := `{"message": "Profile updated successfully"}`

	assert.JSONEq(t, expectedMessage, w.Body.String())

}
