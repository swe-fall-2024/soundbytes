import { Component, ChangeDetectionStrategy, OnInit } from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {FormsModule} from '@angular/forms';
import {CommonModule } from '@angular/common';
import {MatSelectModule} from '@angular/material/select';
import { Profile } from '../../models/profile.model';
import { ProfileService } from '../../profile.service';
import { UserService } from '../../services/signup.component';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-edit-profile',
  imports: [MatSelectModule, CommonModule, MatCardModule, MatInputModule, MatFormFieldModule, FormsModule],
  templateUrl: './edit-profile.component.html',
  styleUrl: './edit-profile.component.css',
  changeDetection: ChangeDetectionStrategy.Default,
})


export class EditProfileComponent implements OnInit {

  user: Profile | null = null;
  userID = String(localStorage.getItem('currentUserEmail')); 

  profile = {
      name: ' ',
      username: '@shiba',
      currentFavType: 'Current Favorite Artist',
      currentFav: 'COIN',
      genres: ['indie','pop', 'hyperpop'],
      topSong: "stupid horse",
      topArtist: "100 gecs",
  };


  constructor(private profileService: ProfileService, private router: Router) {}
  
  ngOnInit(): void {
  this.profileService.getUserProfile(this.userID).subscribe({
    next: (data) => {
      console.log('Data from API:', data); // Log to check if the data looks correct
      this.user = data;
      console.log('User Info:', this.user);
      if (this.user && this.user.topSong) {
        console.log("Cameron's top Song:", this.user.topSong);  // Should log 'Sweet Caroline'
      } else {
        console.log("THE TOP SONG is missing or undefined");
      }

      // After user data is fetched, update profiles
      this.updateProfiles();
    },
    error: (error) => {
      console.error('Error fetching user data:', error);
    }
  });
  }

  updateProfiles() {
    if (this.user) {
      console.log("update profile method is being fired`")
      // Update the profiles array after the user data is available
      this.profile = {
          name: "",
          username: this.user.username,
          currentFavType: 'Current Favorite Artist',
          currentFav: this.user.topArtist, // Now this is updated correctly
          genres: this.user.favGenres,
          topSong: this.user.topSong, // Now this is updated correctly
          topArtist: this.user.topArtist, // Now this is updated correctly
        };
    }
  }

  saveProfile() {
    if (!this.user) {
      console.error('No user data to save');
      return;
    }

    // Prepare the updated user profile data
    const updatedProfile: Profile = {
      userID: this.userID,
      username: this.profile.username,
      password: this.user.password,
      topArtist: this.profile.topArtist,
      topSong: this.profile.topSong,
      favSongs: this.user.favSongs,
      favGenres: this.profile.genres,
      posts: this.user.posts,
      following: this.user.following
    };

    this.profileService.updateUserProfile(this.userID, updatedProfile).subscribe({
      next: (response) => {
        console.log('Profile updated successfully:', response);
        alert("Profile updated successfully:")
        //route to profile page
        this.router.navigate([`/profile`]);
      },
      error: (error) => {
        console.error('Error updating profile:', error);
      }
    });
  }
}