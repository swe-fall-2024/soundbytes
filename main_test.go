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
		Username:  "testuser@gmail.com",
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
		Username: "testuser@gmail.com",
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

	request.Follower = "testuser@gmail.com"
	request.Followee = "juan@ufl.edu"

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

	type User struct {
		UserID    string   `bson:"_id,omitempty" json:"userID"` // Change to match Angular `userID`
		Username  string   `bson:"username" json:"username"`
		Password  string   `json:"password"`
		TopArtist string   `bson:"top_artist" json:"topArtist"`      // Match Angular `topArtist`
		TopSong   string   `bson:"top_song" json:"topSong"`          // Match Angular `topSong`
		FavSongs  []string `bson:"favorite_songs" json:"favSongs"`   // Match Angular `favSongs`
		FavGenres []string `bson:"favorite_genres" json:"favGenres"` // Match Angular `favGenres`
		Posts     []Post   `bson:"posts" json:"posts"`
		Following []string `json:"following"`
	}

	var request struct {
		Username  string   `bson:"username" json:"username"`
		Name      string   `json:"name"`
		TopArtist string   `bson:"top_artist" json:"topArtist"`
		TopSong   string   `bson:"top_song" json:"topSong"`
		FavSongs  []string `bson:"favorite_songs" json:"favSongs"`
		FavGenres []string `bson:"favorite_genres" json:"favGenres"`
	}

	request.Username = "testuser@gmail.com"
	request.Name = "Mr Test"
	request.TopArtist = "testtopartist"
	request.TopSong = "testtopsong"
	request.FavSongs = []string{"We are the Champions"}
	request.FavGenres = []string{"testtopgenre1", "testtopgenre2", "testtopgenre3"}

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

func TestGetProfileHandler(t *testing.T) {
	// Setup the test environment
	setup()

	testUserCollection = testClient.Database("testdb").Collection("users")

	// Define the test username
	testUsername := "testuser@gmail.com"

	// Create a new HTTP GET request with the username as a query parameter
	req := httptest.NewRequest(http.MethodGet, "/getProfile?username="+testUsername, nil)

	w := httptest.NewRecorder()

	// Call the handler
	getProfileHandlerForTesting(w, req, testClient, testUserCollection)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body message
	expectedMessage := `{
	"username":"testuser@gmail.com",
	"name":"Mr Test",
	"topArtist":"testtopartist",
	"topSong":"testtopsong",
	"favSongs":["We are the Champions"],
	"favGenres":["testtopgenre1","testtopgenre2","testtopgenre3"]
	}`

	assert.JSONEq(t, expectedMessage, w.Body.String())

}

func TestRegisterSongHandler(t *testing.T) {
	// Setup the test environment
	setup()

	testUserCollection = testClient.Database("testdb").Collection("songs")

	// Define the test song
	testSong := Song{
		Id:     "4pNiE4LCVV74vfIBaUHm1b",
		Name:   "",
		Artist: "",
	}

	// Convert testSong to JSON
	songJSON, err := json.Marshal(testSong)

	println(string(songJSON))

	if err != nil {
		t.Fatalf("Failed to marshal test song: %v", err)
	}

	// Create a new HTTP POST request with the test song data
	req := httptest.NewRequest(http.MethodPost, "/registerSong", bytes.NewReader(songJSON))

	println(req)

	w := httptest.NewRecorder()

	println(w)

	// Call the register handler

	registerSongHandlerForTesting(w, req, testClient, testUserCollection)

	// Check the response status code

	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body message

	expectedMessage := `{"message": "Song Post registered successfully"}`

	assert.JSONEq(t, expectedMessage, w.Body.String())

	// Verify that the song was actually added to the database

	var storedSong Song

	err = testUserCollection.FindOne(context.TODO(), bson.M{"id": testSong.Id}).Decode(&storedSong)

	if err != nil {

		t.Fatalf("Failed to find the registered song in the database: %v", err)

	}
}

func TestRegisterAlbumHandler(t *testing.T) {
	// Setup the test environment
	setup()

	testUserCollection = testClient.Database("testdb").Collection("albums")

	// Define the test song
	testAlbum := Album{
		Id:          "50o7kf2wLwVmOTVYJOTplm",
		Name:        "",
		Artist:      "",
		Genre:       "",
		ReleaseDate: "",
	}

	// Convert testSong to JSON
	albumJSON, err := json.Marshal(testAlbum)

	println(string(albumJSON))

	if err != nil {
		t.Fatalf("Failed to marshal test album: %v", err)
	}

	// Create a new HTTP POST request with the test song data
	req := httptest.NewRequest(http.MethodPost, "/registerAlbum", bytes.NewReader(albumJSON))

	println(req)

	w := httptest.NewRecorder()

	println(w)

	// Call the register handler

	registerAlbumHandlerForTesting(w, req, testClient, testUserCollection)

	// Check the response status code

	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body message

	expectedMessage := `{"message": "Album Post registered successfully"}`

	assert.JSONEq(t, expectedMessage, w.Body.String())

	// Verify that the song was actually added to the database

	var storedAlbum Song

	err = testUserCollection.FindOne(context.TODO(), bson.M{"id": testAlbum.Id}).Decode(&storedAlbum)

	if err != nil {

		t.Fatalf("Failed to find the registered song in the database: %v", err)

	}
}

func TestUnfollowUserHandler(t *testing.T) {

	// Setup the test environment
	setup()

	testUserCollection = testClient.Database("testdb").Collection("users")

	// Decode request body
	var request struct {
		Follower string `json:"follower"`
		Followee string `json:"followee"`
	}

	request.Follower = "testuser@gmail.com"
	request.Followee = "juan@ufl.edu"

	// Convert testUser to JSON
	requestJSON, err := json.Marshal(request)

	println(string(requestJSON))

	if err != nil {
		t.Fatalf("Failed to marshal test user: %v", err)
	}

	// Create a new HTTP POST request with the test user data
	req := httptest.NewRequest(http.MethodPost, "/unfollow", bytes.NewReader(requestJSON))

	println(req)

	w := httptest.NewRecorder()

	println(w)

	// Call the register handler
	unfollowUserHandlerForTesting(w, req, testClient, testUserCollection)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body message
	expectedMessage := `{"message": "User unfollowed successfully"}`
	assert.JSONEq(t, expectedMessage, w.Body.String())

}

func TestGetFeedHandler(t *testing.T) {

	setup()

	print("It Reached Here Line 439")

	testUsername := "testuser@gmail.com"

	userCollection := testClient.Database(dbName).Collection("users")
	postCollection := testClient.Database(dbName).Collection("posts")

	req := httptest.NewRequest(http.MethodGet, "/getFeed?username="+testUsername, nil)

	w := httptest.NewRecorder()

	getFeedForTesting(w, req, testClient, userCollection, postCollection)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)
}
