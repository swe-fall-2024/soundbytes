import {ChangeDetectionStrategy, Component, OnInit, SimpleChanges} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatChipsModule} from '@angular/material/chips';
import {MatProgressBarModule} from '@angular/material/progress-bar';
import {MatIconModule} from '@angular/material/icon';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatListModule} from '@angular/material/list';
import {MatDividerModule} from '@angular/material/divider';
import {MatToolbarModule} from '@angular/material/toolbar';
import { PlaylistCardComponentComponent } from '../../components/playlist-card/playlist-card.component';
import { ReviewCardComponent } from '../../components/review-card/review-card.component';
import { SongCardComponent } from '../../components/song-card/song-card.component';
import { NgIf, NgFor, CommonModule } from '@angular/common';
import { FriendsComponent } from '../../components/friends/friends.component';
import { BioComponent } from '../../components/bio/bio.component';
import { ProfileService } from '../../profile.service';
import { Profile } from '../../models/profile.model';  // Import the User interface from profile.model.ts

@Component({
  selector: 'app-profile',
  imports: [BioComponent, FriendsComponent, NgIf, NgFor, CommonModule, SongCardComponent, ReviewCardComponent, PlaylistCardComponentComponent, MatToolbarModule, MatDividerModule, MatListModule, MatCardModule, MatButtonModule, MatChipsModule, MatProgressBarModule, MatIconModule, MatGridListModule],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css',
  changeDetection: ChangeDetectionStrategy.Default
})


export class ProfileComponent implements OnInit {
  user: Profile | null = null;
  userID = 'cam123@gmail.com'; // Example user ID (Replace with dynamic value)

  // Static data for posts and friends
  posts = [
    {
      user: 'Shiba Inu',
      profile_img: 'url',
      type: 'favorite-song',
      title: 'MY FAVORITE SONG',
      content: {
        song_title: 'Engagement Party',
        song_url: 'https://open.spotify.com/track/5PYPCxyWltRIyPkhSsnWIk',
        song_embed: "https://open.spotify.com/embed/track/6LxcPUqx6noURdA5qc4BAT?utm_source=generator",
      }
    },
    {
      user: 'Shiba Inu',
      profile_img: 'url',
      type: 'album-review',
      title: 'ALBUM REVIEW',
      content: {
        album_title: "Short n' Sweet",
        review: "Sabrina Carpenter's latest album, Short n' Sweet, released on August 23, 2024, marks her sixth studio endeavor and showcases a refreshingly lighthearted and cheeky approach to pop music. The album has been lauded for its cleverness and effortless execution, setting a high bar for contemporary pop.",
      }
    },
    {
      user: 'Shiba Inu',
      profile_img: 'url',
      type: 'playlist',
      title: 'MY NEW PLAYLIST',
      content: {
        playlist_title: "Study playlist",
        playlist_url: "https://open.spotify.com/playlist/1yJb4XCnM4KfeO2UkMAYnp?si=945e25fe87034d38",
        playlist_embed: "https://open.spotify.com/embed/playlist/1yJb4XCnM4KfeO2UkMAYnp?utm_source=generator",
      }
    },
  ];

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

  constructor(private profileService: ProfileService) {}

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
}
