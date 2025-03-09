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
var testPostCollection *mongo.Collection

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

func TestGetProfileHandler(t *testing.T) {
	// Setup the test environment
	setup()

	testUserCollection = testClient.Database("testdb").Collection("users")

	// Define the test username
	testUsername := "testuser"

	// Create a new HTTP GET request with the username as a query parameter
	req := httptest.NewRequest(http.MethodGet, "/getProfile?username="+testUsername, nil)

	w := httptest.NewRecorder()

	// Call the handler
	getProfileHandlerForTesting(w, req, testClient, testUserCollection)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body message
	expectedMessage := `{"username":"testuser","name":"testname","top_artist":"testtopartist","top_song":"testtopsong","top_genre":"testtopgenre","top_3_genres":["testtopgenre1","testtopgenre2","testtopgenre3"],"all_time_fav_song":"testalltimefavsong","all_time_fav_artist":"testalltimefavartist"}`

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

func TestCreatePostHandler(t *testing.T) {
	// Setup the test environment
	setup()

	testPostCollection = testClient.Database("testdb").Collection("posts")

	testPost := Post{
		PostID:    "1234567890",
		PostType:  "Facebook Post",
		Link:      "www.facebook.com",
		Title:     "Grandma's Post",
		Content:   "I visited Junior last week. He was nice.",
		LikeCount: 500000,
	}

	//Convert testPost to JSON
	postJSON, err := json.Marshal(testPost)

	println(string(postJSON))

	if err != nil {
		t.Fatalf("Failed to marshal post: %v", err)
	}

	// Create a new HTTP POST request with the test post data
	req := httptest.NewRequest(http.MethodPost, "/addPost", bytes.NewReader(postJSON))
	req.Header.Set("Content-Type", "application/json")

	println(req)

	w := httptest.NewRecorder()

	println(w)

	// Call the register handler

	addPostHandlerForTesting(w, req, testClient, testPostCollection)

	// Check the response status code

	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body message

	expectedMessage := `{"message": "Post created successfully"}`

	assert.JSONEq(t, expectedMessage, w.Body.String())

	// Verify that the song was actually added to the database

	var storedPost Post

	err = testPostCollection.FindOne(context.TODO(), bson.M{"_id": testPost.PostID}).Decode(&storedPost)

	if err != nil {

		t.Fatalf("Failed to find the post in the database: %v", err)

	}

}

func TestDeletePostHandler(t *testing.T) {
	// Setup the test environment
	setup()

	testPostCollection = testClient.Database("testdb").Collection("posts")

	// Create a test post
	testPost := Post{
		PostID:    "0987654321",
		PostType:  "MySpace Post",
		Link:      "www.myspace.com",
		Title:     "Opa's Post",
		Content:   "I did not visit Junior last week. It was also nice.",
		LikeCount: 500000,
	}

	// Insert test post into the database
	_, err := testPostCollection.InsertOne(context.TODO(), testPost)
	if err != nil {
		t.Fatalf("Failed to insert post into database: %v", err)
	}

	// Verify that the post was added successfully
	var storedPost Post
	err = testPostCollection.FindOne(context.TODO(), bson.M{"_id": testPost.PostID}).Decode(&storedPost)
	if err != nil {
		t.Fatalf("Failed to find the post in the database: %v", err)
	}

	// Create a DELETE request
	req := httptest.NewRequest(http.MethodDelete, "/deletePost", bytes.NewBufferString(`{"postID": "0987654321"}`))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Call the delete handler
	deletePostHandlerForTesting(w, req, testClient, testPostCollection)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body message
	expectedMessage := `{"message": "Post deleted successfully"}`
	assert.JSONEq(t, expectedMessage, w.Body.String())

	// Verify that the post no longer exists in the database
	err = testPostCollection.FindOne(context.TODO(), bson.M{"_id": testPost.PostID}).Decode(&storedPost)
	if err == nil {
		t.Fatalf("Post was not deleted from the database")
	}
}

func TestUpdatePostHandler(t *testing.T) {
	// Setup the test environment
	setup()

	testPostCollection = testClient.Database("testdb").Collection("posts")

	// Create a test post
	testPost := Post{
		PostID:    "5678901234",
		PostType:  "Reddit Post",
		Link:      "www.reddit.com",
		Title:     "Uncle's Post",
		Content:   "I argued with Junior last week. It was interesting.",
		LikeCount: 2500,
	}

	// Insert test post into the database
	_, err := testPostCollection.InsertOne(context.TODO(), testPost)
	if err != nil {
		t.Fatalf("Failed to insert post into database: %v", err)
	}

	// Verify that the post was added successfully
	var storedPost Post
	err = testPostCollection.FindOne(context.TODO(), bson.M{"_id": testPost.PostID}).Decode(&storedPost)
	if err != nil {
		t.Fatalf("Failed to find the post in the database: %v", err)
	}

	// Create an UPDATE request with new title and content
	updateData := `{"postID": "5678901234", "title": "Updated Uncle's Post", "content": "I changed my mind about last week."}`
	req := httptest.NewRequest(http.MethodPut, "/updatePost", bytes.NewBufferString(updateData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Call the update handler
	updatePostHandlerForTesting(w, req, testClient, testPostCollection)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body message
	expectedMessage := `{"message": "Post updated successfully"}`
	assert.JSONEq(t, expectedMessage, w.Body.String())

	// Verify that the post was updated correctly in the database
	err = testPostCollection.FindOne(context.TODO(), bson.M{"_id": testPost.PostID}).Decode(&storedPost)
	if err != nil {
		t.Fatalf("Failed to find the updated post in the database: %v", err)
	}

	assert.Equal(t, "Updated Uncle's Post", storedPost.Title)
	assert.Equal(t, "I changed my mind about last week.", storedPost.Content)
}

func TestUpdateNonExistentPost(t *testing.T) {
	setup()

	testPostCollection = testClient.Database("testdb").Collection("posts")

	// Create an UPDATE request for a non-existent post
	updateData := `{"postID": "0000000000", "title": "Nonexistent Post", "content": "This shouldn't exist."}`
	req := httptest.NewRequest(http.MethodPut, "/updatePost", bytes.NewBufferString(updateData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Call the update handler
	updatePostHandlerForTesting(w, req, testClient, testPostCollection)

	// Check the response status code
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Check the response body message
	expectedMessage := "Post not found\n"
	assert.Equal(t, expectedMessage, w.Body.String())
}

func TestUpdatePostWithNoFields(t *testing.T) {
	setup()

	testPostCollection = testClient.Database("testdb").Collection("posts")

	// Create a test post
	testPost := Post{
		PostID:    "1231231234",
		PostType:  "Twitter Post",
		Link:      "www.twitter.com",
		Title:     "Original Tweet",
		Content:   "This is an unedited tweet.",
		LikeCount: 200,
	}

	_, err := testPostCollection.InsertOne(context.TODO(), testPost)
	if err != nil {
		t.Fatalf("Failed to insert post into database: %v", err)
	}

	// Create an UPDATE request with an empty body
	updateData := `{"postID": "1231231234"}`
	req := httptest.NewRequest(http.MethodPut, "/updatePost", bytes.NewBufferString(updateData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Call the update handler
	updatePostHandlerForTesting(w, req, testClient, testPostCollection)

	// Check the response status code
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// Check the response body message
	expectedMessage := "No fields to update\n"
	assert.Equal(t, expectedMessage, w.Body.String())
}

func TestDeleteNonExistentPost(t *testing.T) {
	setup()

	testPostCollection = testClient.Database("testdb").Collection("posts")

	// Create a DELETE request for a non-existent post
	deleteData := `{"postID": "9999999999"}`
	req := httptest.NewRequest(http.MethodDelete, "/deletePost", bytes.NewBufferString(deleteData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Call the delete handler
	deletePostHandlerForTesting(w, req, testClient, testPostCollection)

	// Check the response status code
	assert.Equal(t, http.StatusNotFound, w.Code)

	// Check the response body message
	expectedMessage := "Post not found\n"
	assert.Equal(t, expectedMessage, w.Body.String())
}

func TestAddDuplicatePost(t *testing.T) {
	setup()

	testPostCollection = testClient.Database("testdb").Collection("posts")

	testPost := Post{
		PostID:    "2222222222",
		PostType:  "LinkedIn Post",
		Link:      "www.linkedin.com",
		Title:     "Networking Advice",
		Content:   "Here’s how you can grow your network!",
		LikeCount: 5000,
	}

	_, err := testPostCollection.InsertOne(context.TODO(), testPost)
	if err != nil {
		t.Fatalf("Failed to insert post into database: %v", err)
	}

	// Convert testPost to JSON
	postJSON, err := json.Marshal(testPost)
	if err != nil {
		t.Fatalf("Failed to marshal post: %v", err)
	}

	// Create a new HTTP POST request with the duplicate post data
	req := httptest.NewRequest(http.MethodPost, "/addPost", bytes.NewReader(postJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Call the add post handler again
	addPostHandlerForTesting(w, req, testClient, testPostCollection)

	// Check for duplicate key error (assuming unique index on _id)
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	expectedMessage := "Error posting post\n"
	assert.Equal(t, expectedMessage, w.Body.String())
}

func TestDeleteExistingPost(t *testing.T) {
	setup()

	testPostCollection = testClient.Database("testdb").Collection("posts")

	// Create a test post
	testPost := Post{
		PostID:    "3333333333",
		PostType:  "Instagram Post",
		Link:      "www.instagram.com",
		Title:     "Vacation Pictures",
		Content:   "Look at these beautiful beaches!",
		LikeCount: 10000,
	}

	_, err := testPostCollection.InsertOne(context.TODO(), testPost)
	if err != nil {
		t.Fatalf("Failed to insert post into database: %v", err)
	}

	// Verify that the post was added successfully
	var storedPost Post
	err = testPostCollection.FindOne(context.TODO(), bson.M{"_id": testPost.PostID}).Decode(&storedPost)
	if err != nil {
		t.Fatalf("Failed to find the post in the database: %v", err)
	}

	// Create a DELETE request
	deleteData := `{"postID": "3333333333"}`
	req := httptest.NewRequest(http.MethodDelete, "/deletePost", bytes.NewBufferString(deleteData))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	// Call the delete handler
	deletePostHandlerForTesting(w, req, testClient, testPostCollection)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body message
	expectedMessage := `{"message": "Post deleted successfully"}`
	assert.JSONEq(t, expectedMessage, w.Body.String())

	// Verify that the post no longer exists in the database
	err = testPostCollection.FindOne(context.TODO(), bson.M{"_id": testPost.PostID}).Decode(&storedPost)
	if err == nil {
		t.Fatalf("Post was not deleted from the database")
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

	request.Follower = "testuser"
	request.Followee = "testuser234"

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
