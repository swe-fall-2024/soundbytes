# Sprint 4 Report

## Work Completed in Sprint 4

### Backend  
During Sprint 4, the backend was enhanced with new functionalities that elevate user interaction and social features:  

- **Like Post**: Added functionality to update user profile information such as favorite songs, artists, and genres.  
- **Search For Users To Follow**: Enabled users to search for and follow others, fostering a more connected user experience.   

### Frontend  
- **Profile Pictures**: Added functionality to assign different user profile picture to each user.
- **Post Creation Formatting**: 
- **Updated Frontend Tests**: 

## Youtube Demo

## Unit Tests for Frontend
Below is the list of unit tests implemented for frontend functionality, implemented using jasmine and karma:

### AppComponent
- ensures that the component is created and launched successfully
- validates that the title in the component is correct
- checks that the navbar component is rendered within it

### SignupComponent
- Ensures that the signup page is created and launched successfully

### LoginComponent
- Ensures that the login page is created and launched successfully

### PostCreationComponent
-  Ensures that post creation page has no errors and is deployed fully

### ProfileComponent
- Verifies that the profile is fully functional

### FeedComponent
- Ensures that the feed component works and is created

### EditProfileComponent
- Ensures that the edit profile component is fully correct and can be launched
  
### FriendProfileComponent
-  Ensures that the friend profile page has no errors and is created

### BioComponent
- Verifies that the bio component is correct and has no errors on creation

### FriendsComponent
- Ensures that the friends component is created and can be used successfully

### NavBarComponent
- Ensures that the navbar is created without errors

### PlaylistCardComponent
- Ensures that the playlist card is successfully created to be displayed in other components
  
### ReviewCardComponent
- Ensures that the review card is successfully created to be displayed in other components
  
### SongCardComponent
- Ensures that the navbar is created without errors

### SearchComponent
- Ensures that the review card is successfully created to be displayed in other components
  
### ProfileService
- Ensures that the profile service that updates a user's profile is functioning properly

## Cypress Test for Frontend

### Test1: Login form
- Navigates to login page
- Fills in the input for email and password
- Clicks the button to submit
- Checks for an alert to say 'Login Successful'

## Backend Unit Tests  

### TestRegisterHandler  
- Tests user registration and ensures the password is hashed.  
- Confirms response code is **201 Created** and success message is returned.  
- Verifies user was added to MongoDB.  

### TestLoginHandler  
- Tests user login functionality.  
- Verifies **200 OK** status is returned upon successful login.  

### TestFollowUserHandler  
- Tests user follow functionality.  
- Asserts successful addition of a follow relationship and proper success response.  

### TestUnfollowUserHandler  
- Tests user unfollow functionality.  
- Verifies successful removal of a follow relationship and correct status response.  

### TestSetUpProfileHandler  
- Verifies profile setup allows storing and updating user’s favorite songs, artists, genres, and name.  
- Checks for success response and proper data persistence.  

### TestGetProfileHandler  
- Retrieves a user's profile using the username.  
- Ensures the correct data fields are returned as expected.  

### TestRegisterSongHandler  
- Tests adding a song post to the database.  
- Asserts response is **201 Created** and post is stored in the collection.  

### TestRegisterAlbumHandler  
- Tests adding an album post to the database.  
- Confirms expected response and that the album is stored correctly.  

### TestGetFeedHandler  
- Tests retrieving posts from followed users for a given user.  
- Confirms the feed is returned successfully with **HTTP 200 OK**.

### TestLikePost
- Tests the backend's ability to handle a POST request to like a specific post.  
- Simulates sending a like request for a post with ID `"92842"`.  
- Verifies that the server responds with **HTTP 200 OK**.  
- Confirms the response body contains the message: `{"message": "Post liked successfully"}`.
- 
# Backend API Documentation

The backend API includes the following endpoints:  

## Authentication  

### `POST /register`  
Registers a new user.  

### `POST /login`  
Authenticates a user.  

## User Management  

### `POST /follow`  
Allows a user to follow another user.  

### `POST /unfollow`  
Allows a user to unfollow another user.  

### `PUT /setUpProfile`  
Creates or updates a user profile.  

### `GET /getProfile`  
Retrieves user profile information.  

### `PUT /editProfile`  
Updates a user’s profile with new information.  

## Posts Management  

### `POST /registerSong`  
Registers a song post.  

### `POST /addPost`  
Adds a post.  

### `PUT /updatePost/{post_id}`  
Updates a post.  

### `DELETE /deletePost/{post_id}`  
Deletes a post.  

### `GET /getPost/{post_id}`  
Gets a post.  

## User Feed  

### `GET /feed?username={username}`  
Retrieves posts from users that the specified user is following.  

**Query Parameter:**  
- `username` – the username of the current user.  

