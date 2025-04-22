import {ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatChipsModule} from '@angular/material/chips';
import {MatProgressBarModule} from '@angular/material/progress-bar';
import {MatIconModule} from '@angular/material/icon';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatListModule} from '@angular/material/list';
import {MatDividerModule} from '@angular/material/divider';
import {MatToolbarModule} from '@angular/material/toolbar';
import { BioComponent } from '../../components/bio/bio.component';
import { NgIf, NgFor, CommonModule } from '@angular/common';
import { SongCardComponent } from '../../components/song-card/song-card.component';
import { ReviewCardComponent } from '../../components/review-card/review-card.component';
import { PlaylistCardComponent } from '../../components/playlist-card/playlist-card.component';
import { ActivatedRoute } from '@angular/router';
import { ProfileService } from '../../profile.service';
import { Profile } from '../../models/profile.model';  // Import the User interface from profile.model.ts
import { HttpClient } from '@angular/common/http';
import { NavbarComponent } from '../../components/navbar/navbar.component';


@Component({
  selector: 'app-friend-profile',
  imports: [NavbarComponent, BioComponent, NgIf, NgFor, CommonModule, SongCardComponent, ReviewCardComponent, PlaylistCardComponent, MatToolbarModule, MatDividerModule, MatListModule, MatCardModule, MatButtonModule, MatChipsModule, MatProgressBarModule, MatIconModule, MatGridListModule],
  templateUrl: './friend-profile.component.html',
  styleUrl: './friend-profile.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class FriendProfileComponent implements OnInit {
  user: Profile | null = null;
  userID: string | null = null;
    
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
  randomIndex = 10;

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
        name: 'Katie',
        username: 'katie123@gmail.com',
        favTypeCurrent: 'Current Favorite Artist',
        favCurrent: 'Sabrina Carpenter',
        genres: ['rock', 'pop', 'hyperpop'],
        topSong: 'Bad Reviews',
        topArtist: 'Sabrina Carpenter',
        pic: this.profileImages.at(this.randomIndex)
      }
    ];
  
  
    constructor(private route: ActivatedRoute, private profileService: ProfileService, private http: HttpClient, private cdr: ChangeDetectorRef) {}
  
    ngOnInit(): void {
      this.userID = String(this.route.snapshot.paramMap.get('id')) //"katie123@gmail.com" //

      this.profileService.getUserProfile(this.userID).subscribe({
        next: (data) => {
          console.log('Data from API..:', data); // Log to check if the data looks correct
          this.user = data;
          console.log('User Info:', this.user);
          if (this.user && this.user.topSong) {
            console.log("Cam's top Song:", this.user.topSong);  // Should log 'Sweet Caroline'
          } else {
            console.log("topSong is missing or undefined");
          }

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
  
      // Fetch posts data
      this.fetchPosts();
  
    }
  
    updateProfiles() {
      if (this.user) {
        // Update the profiles array after the user data is available
        this.profiles = [
          {
            name: this.user.name,
            username: this.user.username,
            favTypeCurrent: this.user.favTypeCurrent,
            favCurrent: this.user.favCurrent, // Now this is updated correctly
            genres: this.user.favGenres,
            topSong: this.user.topSong, // Now this is updated correctly
            topArtist: this.user.topArtist, // Now this is updated correctly
            pic: this.im,
          }
        ];
      }
    }
    
    fetchPosts() {
      if (this.userID = "Katie") {
        this.userID = "katie123@gmail.com"
      }
  
      this.http.get<any[]>(`http://127.0.0.1:4201/getPosts/${this.userID}`).subscribe(
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
  
    me = false;
    following = true;
  }
