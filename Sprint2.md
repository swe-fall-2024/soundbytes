# Sprint 2 Report

## Work Completed in Sprint 2
During Sprint 2, we successfully integrated the front-end and back-end of our application. Additionally, we implemented several key features in the backend, including:
- A function to follow a friend within the app.
- A function to create a user profile and store it in the MongoDB database.
- Unit tests for each of our previously created backend functions.

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
- **PUT /setUpProfile**: Creates or updates a user profile.
- **GET /getProfile**: Retrieves user profile information.
- **POST /registerSong**: Registers a song post.

Each API endpoint is tested using unit tests to ensure functionality and reliability.

