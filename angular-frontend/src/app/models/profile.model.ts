interface Post {
  user: string;
  profile_img: string;
  type: string;
  title: string;
  content: any;  // Adjust based on the actual structure of content (e.g., if it's a specific type)
}

export interface Profile {
  userID: string;  // Corresponding to the Go field `UserID`
  name: string;
  username: string;
  password: string;
  topArtist: string;
  topSong: string;
  favSongs: string[]; // List of favorite songs
  favGenres: string[]; // List of favorite genres
  posts: Post[]; // List of posts (you can adjust the `Post` interface as needed)
  following: string[]; // List of usernames the user is following
  pic: string;
}
