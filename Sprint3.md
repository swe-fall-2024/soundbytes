# Sprint 3 Report

## Work Completed in Sprint 3

### Backend  
During Sprint 3, the backend was enhanced with new functionalities that elevate user interaction and social features:  

- **Edit Profile**: Added functionality to update user profile information such as favorite songs, artists, and genres.  
- **Review Posting**: Users can now post reviews (songs, albums, playlists), which are saved in the backend.  
- **User Feed**: Developed the `getFeed` endpoint to serve posts from users that the current user is following.  

### Frontend  
On the frontend, we improved interactivity, personalization, and routing:  

- **User State Management**: Dynamically stored the logged-in user and made the data accessible throughout the site.  
- **Post Songs Page UI**: Improved form design and made the experience more intuitive for users.  
- **Routing and Navigation**: Refined the navbar and page routing for smoother and more predictable navigation.

## Youtube Demo
https://youtu.be/93jt3K3rwqQ

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

