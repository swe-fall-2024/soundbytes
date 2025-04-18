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


  profileImages: string[] = [
    '/1.jpg',
    '/2.jpg',
    '/3.jpg',
    '/4.jpg',
    '/5.jpg',
    '/6.jpg',
    '/7.jpg',
    '/8.jpg',
    '/9.jpg',
    '/10.jpg',
    '/11.jpg',
    '/12.jpg'
  ];


  im!: string;
  randomIndex = 11;

  profile = {
      name: ' ',
      username: '@shiba',
      currentFavType: 'Current Favorite Artist',
      currentFav: 'COIN',
      genres: ['indie','pop', 'hyperpop'],
      topSong: "stupid horse",
      topArtist: "100 gecs",
      pic: this.profileImages.at(this.randomIndex)
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
      }
      else {
        console.log("THE TOP SONG is missing or undefined");
      }
      console.log("HERE");

      if(localStorage.getItem('count') != null) {
        this.randomIndex = parseInt(localStorage.getItem('count')!);
        if(this.randomIndex > 11) {
          localStorage.setItem('count', '0');
          this.randomIndex = parseInt(localStorage.getItem('count')!);
        }
      }

      if (this.user.pic == ""){
        this.im = this.profileImages.at(this.randomIndex)!;
        console.log(this.im);
      } 
      else {
        this.im = this.user.pic;
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
          name: this.user.name,
          username: this.user.username,
          currentFavType: 'Current Favorite Artist',
          currentFav: this.user.topArtist, // Now this is updated correctly
          genres: this.user.favGenres,
          topSong: this.user.topSong, // Now this is updated correctly
          topArtist: this.user.topArtist, // Now this is updated correctly
          pic: this.im
        };
        console.log('this is profile pic: ', this.profile.pic);
        console.log(this.im);
    }
  }

  saveProfile() {
    console.log("in save profile")
    if (!this.user) {
      console.error('No user data to save');
      return;
    }
    
    // Prepare the updated user profile data
    const updatedProfile: Profile = {
      userID: this.userID,
      username: this.profile.username,
      name: this.profile.name,
      password: this.user.password,
      topArtist: this.profile.topArtist,
      topSong: this.profile.topSong,
      favSongs: this.user.favSongs,
      favGenres: this.profile.genres,
      posts: this.user.posts,
      following: this.user.following,
      pic: this.im
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
