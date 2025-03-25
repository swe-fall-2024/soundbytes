import {ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit, SimpleChanges} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatChipsModule} from '@angular/material/chips';
import {MatProgressBarModule} from '@angular/material/progress-bar';
import {MatIconModule} from '@angular/material/icon';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatListModule} from '@angular/material/list';
import {MatDividerModule} from '@angular/material/divider';
import {MatToolbarModule} from '@angular/material/toolbar';
import { PlaylistCardComponent } from '../../components/playlist-card/playlist-card.component';
import { ReviewCardComponent } from '../../components/review-card/review-card.component';
import { SongCardComponent } from '../../components/song-card/song-card.component';
import { NgIf, NgFor, CommonModule } from '@angular/common';
import { FriendsComponent } from '../../components/friends/friends.component';
import { BioComponent } from '../../components/bio/bio.component';
import { ProfileService } from '../../profile.service';
import { Profile } from '../../models/profile.model';  // Import the User interface from profile.model.ts
import { HttpClient } from '@angular/common/http';
import { NavbarComponent } from '../../components/navbar/navbar.component';

@Component({
  selector: 'app-profile',
  imports: [BioComponent, FriendsComponent, NavbarComponent, NgIf, NgFor, CommonModule, SongCardComponent, ReviewCardComponent, PlaylistCardComponent, MatToolbarModule, MatDividerModule, MatListModule, MatCardModule, MatButtonModule, MatChipsModule, MatProgressBarModule, MatIconModule, MatGridListModule],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css',
  changeDetection: ChangeDetectionStrategy.Default
})


export class ProfileComponent implements OnInit {
  user: Profile | null = null;
  userID = 'cam123@gmail.com'; // Example user ID (Replace with dynamic value)

  // Static data for posts and friends
  posts: any[] = [];

  // Static friends list
  friends = [
    { name: 'Katie' },
    { name: 'Mary' }
  ];

  // Initial profile placeholder, it will be updated once user data is fetched
  profiles = [
    {
      name: 'Shibaaaaaa Inu',
      username: '@shiba',
      currentFavType: 'Current Favorite Artist',
      currentFav: '',
      genres: ['indie', 'pop', 'hyperpop'],
      topSong: '',
      topArtist: '',
    }
  ];


  constructor(private profileService: ProfileService, private http: HttpClient, private cdr: ChangeDetectorRef) {}

  ngOnInit(): void {

    this.profileService.getUserProfile(this.userID).subscribe({
      next: (data) => {
        console.log('Data from API:', data); // Log to check if the data looks correct
        this.user = data;
        console.log('User Info:', this.user);
        if (this.user && this.user.topSong) {
          console.log("Cam's top Song:", this.user.topSong);  // Should log 'Sweet Caroline'
        } else {
          console.log("topSong is missing or undefined");
        }

        // After user data is fetched, update profiles
        this.updateProfiles();
        this.updateFriends();
      },
      error: (error) => {
        console.error('Error fetching user data:', error);
      }
    });

    // Fetch posts data
    this.fetchPosts();

  }

  updateProfiles() {
    if (this.user) {
      // Update the profiles array after the user data is available
      this.profiles = [
        {
          name: 'Shibaaaaaa Inu',
          username: this.user.username,
          currentFavType: 'Current Favorite Artist',
          currentFav: this.user.topArtist, // Now this is updated correctly
          genres: this.user.favGenres,
          topSong: this.user.topSong, // Now this is updated correctly
          topArtist: this.user.topArtist, // Now this is updated correctly
        }
      ];
    }
  }
  updateFriends() {
    if (this.user) {
      console.log("friends: ", this.friends)
      this.friends = this.user.following?.map((name: string) => ({ name })) || [];    }
      console.log("friends: ", this.friends)
    }
  fetchPosts() {

    this.http.get<any[]>('http://127.0.0.1:4201/getPosts/cam123@gmail.com').subscribe(
      (data) => {
        this.posts = data; // Store the posts data in the component
        //alert(`${this.posts.length} posts fetched`); 
        this.cdr.markForCheck();
      },
      (error) => {
        console.error('Error fetching posts', error);
      }
    );
  
  }

  me = true;
  following = false;

}
