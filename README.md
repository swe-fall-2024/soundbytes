# Sound Bytes

Project Group Members: Cameron Santiago, Katherine Hartley, Juan Veliz, Mary Hanson

# Description

Popular music apps allow a user to listen to music and sometimes follow friends, but users cannot share music, artists, playlists, or concert tickets within these platforms. So, this application creates a more social-centered music experience. Sound Bytes is a music sharing platform that users can use to discover new music or connect with their friends. Additionally, It integrates concerts where users can flag their favorite artists and plan the event with a friend.

# APIs
- songKick
- spotify
- genius
- sound cloud
- apple music
- amazon music
- ticketmaster

# User Manual

Soundbytes is an angular-go web application. It runs on localhost in any browser, but since testing utilized chrome, chrome is the preferred platform.

## Access the Web App

1) git clone this repository, soundbytes, into a desired directory through the terminal
2) cd into this repository from the terminal
3) git checkout merged-frontend (for the most up-to-date functionality)
   
## Run SoundBytes

### Required Dependencies:
- Go: ensure your system has go installed to run the backend api
- mongodb: ensure mongodb is installed and running for the backend's connection
- node, npm, & ng: to install all necessary dependencies npm install will need to be called

Run Frontend
1) Navigate to the soundbytes directory and then to angular-frontend in the terminal
2) Once cded into angular-frontend, run npm install
3) With all dependencies installed, run ng serve
4) The above command should provide you with a link to the application running on your localhost at port 4200
5) Navigate to http://localhost:4200/ to view the soundbytes application!

**The defualt port for angular applications is 4200, but it may be different so follow the link specifically provided by ng serve**

Run Backend
1) In a separate terminal tab or window, navigate to the soundbytes directory
2) Once cded into soundbytes, run go run main.go
3) This should automatically connect to the mongodb connection (ensure mongodb is running and open to connections)
4) With that, the app is ready to be used! **Note that both backend and frontend commands must be running at the same time**

Userflow for Application
1) Start at the root route / which is the login page
2) If not signed up, navigate to the signup page through the link on login
3) If signing up, user is automatically taken to /edit-profile to set up profile
4) Once profile is setup or logged in, the application routes to profile
5) From profile, user can access post creation, feed, profile, or search by the navbar 
