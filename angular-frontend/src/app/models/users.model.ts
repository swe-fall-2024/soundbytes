interface Post {
    user: string;
    profile_img: string;
    type: string;
    title: string;
    content: any;  // Adjust based on the actual structure of content (e.g., if it's a specific type)
}  

export interface User {
    userID: string;      // Matches Go struct 'UserID'
    username: string;    // Matches Go struct 'Username'
    password: string;    // Matches Go struct 'Password'
    topArtist: string;   // Matches Go struct 'TopArtist'
    topSong: string;     // Matches Go struct 'TopSong'
    favSongs: string[];  // Matches Go struct 'FavSongs' (array of strings)
    favGenres: string[]; // Matches Go struct 'FavGenres' (array of strings)
    posts: Post[];       // Assuming 'Post' is another interface or type in Angular
    following: string[]; // Matches Go struct 'Following'
  }