**User Stories** 

- As a music streaming service user, I want to post individual songs or playlists as my favorites, so I can share them with my friends and they can easily access my recommendations.

- As a concert enthusiast, I want to chat with my friends about upcoming concerts in our area, so we can plan to attend together.

- As a music lover and Letterboxd user, I want to be able to post ratings and reviews of albums.

- As a music artist, I want my fans to see my concert dates and new music, so they can stay up-to-date on my content.

- As a long term fan, I do not want to miss when artists I listen to release new albums and songs.

- As a user, I want to securely log in or sign up using my credentials so I can safely access my personalized dashboard and protected features without unauthorized access. 

**Issues the team planned to address**

*FrontEnd*

- Wireframes: Develop and finalize wireframes to outline the structure and layout for key pages.

- Login Page: Implement the login page to allow users to securely access their accounts (pending final confirmation).

- Profile Visualization: Design and integrate an improved user profile visualization to enhance the clarity and accessibility of user information.

- Feed Visualization: Optimize the feed layout for a more intuitive and engaging display of content.

*BackEnd*

- Basic Backend API Setup: Establish a foundational backend API to handle connections with the frontend.

- Mux Router: Set up the Mux router to manage navigation routing across different endpoints.

- RESTful API Structure: Ensure the API follows RESTful conventions with routes like /users, /posts, /feed, and others to handle necessary data retrieval and interaction.

- Route Handlers: Plan and implement route handlers for key functionalities such as user profile data, song/concert sharing, and feed retrieval.

- Firebase Authentication: Implement email/password authentication as well as support for social logins through Firebase.

- Secure API Endpoints: Protect API endpoints by validating Firebase tokens to ensure authorized access.

- Spotify API Setup: Set up a system to interact with the Spotify API to fetch user-specific music data.

- OAuth Authentication: Implement OAuth for user authentication to access their specific Spotify data.

- Fetch User’s Current Track: Enable the backend to retrieve the currently playing track of the user.

- Sharing Functionality: Allow users to share songs, albums, and playlists with others.

- Song, Album, Playlist Info: Retrieve detailed information about songs, albums, and playlists for sharing and display purposes.

**Backend Video Link**

https://youtu.be/vRQUZ1wKv5c

**Backend: Items Successfully Completed**

- Basic Backend API Setup: Establish a foundational backend API to handle connections with the frontend.

- Initial Routing Logic Developed: Developed the initial logic to handle almost all navigation routing, using Gorilla Mux.

- Basic Database Implementation: Set up a basic database structure to account for users, albums, songs ,artists , friends and how to interconnect them.

- Implement a way to retrieve song, album or playlist Info from the Spotify API for sharing and display purposes.

- Basic Login Auth Setup: Basic Login Auth was set up, but we're still working on getting email and or google auth up and running as well. 

**Backend: Which Issues did not get completed and Why** 

Firebase Authentication: Implement email/password authentication as well as support for social logins through Firebase.

- Why: Did not have time to get to this issue, we wanted to get the basic authentification through MongoDB finished before we looked at Firebase Email Based Authentification. 

Secure API Endpoints: Protect API endpoints by validating Firebase tokens to ensure authorized access.

- Why: Have not gotten around to the Firebase Auth Ticket so this will naturally follow suit after that other ticket gets finished. 

Fetch User’s Current Track: Enable the backend to retrieve the currently playing track of the user.

- Why: Currently working the fine tunings of friends/followers logic , once that's done then this will get priority. 

**Frontend Video Link**

https://www.youtube.com/watch?v=KyJ_jJHyZSE

**Frontend: Items Successfully Completed**

- Wireframes: Developed and finalized wireframes to outline the structure and layout for key pages, including login, signup, profile, post creation, and feed.

- Profile Visualization: Designed and integrated an improved user profile visualization to enhance the clarity and accessibility of user information.

- Feed Visualization: Created and structured basic feed page. Optimized the feed layout for a more intuitive and engaging display of content.

**Frontend: Which Issues did not get completed and Why** 

Login and sign up pages

- Why: Prioritized website core functionalities and plan to complete at the start of the next sprint.









