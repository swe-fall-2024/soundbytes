# Sprint 2 Report

## Work Completed in Sprint 2
During Sprint 2, we successfully integrated the front-end and back-end of our application. Additionally, we implemented several key features in the backend, including:
- A function to follow a friend within the app.
- A function to unfollow a friend within the app.
- A function to create a user profile and store it in the MongoDB database.
- Unit tests for each of our previously created backend functions.

And, for frontend, the key features include:
- Functional navigation to different pages through the navbar.
- Friend profile page with follow button.
- An edit profile page to input user details.
- A post creation page that dynamically changes the inputs based on post type.
- Unit tests for each of our frontend components.
- A cypress test that tests the login component by filling in form items and checking feedback from backend.

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

## Cypress Test for Frontend

### Test1: Login form
- Navigates to login page
- Fills in the input for email and password
- Clicks the button to submit
- Checks for an alert to say 'Login Successful'
  
## Unit Tests for Backend
Below is the list of unit tests implemented for backend functionality:

### TestRegisterHandler
- Validates user registration by ensuring the new user is added to the database.
- Confirms that the password is securely hashed.
- Checks response status and success message.

### TestLoginHandler
- Verifies successful user login.
- Ensures that the system correctly authenticates valid credentials.

### TestFollowUserHandler
- Ensures that a user can follow another user successfully.
- Checks response status and success message.

### TestUnfollowUserHandler
-  Ensures that a user can unfollow another user successfully.
-  Checks response status and success message.

### TestSetUpProfileHandler
- Verifies that a user's profile can be created and updated.
- Ensures all profile fields are stored correctly.
- Checks response status and success message.

### TestGetProfileHandler
- Retrieves user profile details based on the username.
- Verifies the correctness of returned profile information.

### TestRegisterSongHandler
- Ensures that a song post can be registered successfully.
- Checks response status and success message.

## Backend API Documentation
The backend API includes the following endpoints:

- **POST /register**: Registers a new user.
- **POST /login**: Authenticates a user.
- **POST /follow**: Allows a user to follow another user.
- **POST /unfollow**: Allows a user to unfollow another user.
- **PUT /setUpProfile**: Creates or updates a user profile.
- **GET /getProfile**: Retrieves user profile information.
- **POST /registerSong**: Registers a song post.

Each API endpoint is tested using unit tests to ensure functionality and reliability.

