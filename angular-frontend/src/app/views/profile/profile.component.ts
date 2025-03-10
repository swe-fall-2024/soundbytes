import { ChangeDetectionStrategy, ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatChipsModule } from '@angular/material/chips';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatIconModule } from '@angular/material/icon';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatListModule } from '@angular/material/list';
import { MatDividerModule } from '@angular/material/divider';
import { MatToolbarModule } from '@angular/material/toolbar';
import { PlaylistCardComponentComponent } from '../../components/playlist-card/playlist-card.component';
import { ReviewCardComponent } from '../../components/review-card/review-card.component';
import { SongCardComponent } from '../../components/song-card/song-card.component';
import { NgIf, NgFor, CommonModule } from '@angular/common';
import { FriendsComponent } from '../../components/friends/friends.component';
import { BioComponent } from '../../components/bio/bio.component';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-profile',
  imports: [BioComponent, FriendsComponent, NgIf, NgFor, CommonModule, SongCardComponent, ReviewCardComponent, PlaylistCardComponentComponent, MatToolbarModule, MatDividerModule, MatListModule, MatCardModule, MatButtonModule, MatChipsModule, MatProgressBarModule, MatIconModule, MatGridListModule],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ProfileComponent implements OnInit {

  posts: any[] = [];
  friends = [
    { name: 'Katie' },
    { name: 'Mary' }
  ];
  profiles  = [{
    name: 'Shiba Inu',
    username: '@shiba',
    currentFavType: 'Current Favorite Artist',
    currentFav: 'COIN',
    genres: {one:'indie',two:'pop',three:'hyperpop'},
    topSong: "stupid horse",
    topArtist: "100 gecs",
  }
]

  me = true;
  following = false;

  constructor(private http: HttpClient, private cdr: ChangeDetectorRef) {}

  ngOnInit() {
    this.fetchPosts();
    //this.fetchProfile(); // Fetch posts when the component initializes
  }

  fetchProfile() {

    this.http.get<any[]>('http://127.0.0.1:4201/getProfile/testuser').subscribe(
      (data) => {
        this.profiles = data; 
        //alert(`${data.length} profiles fetched`); 
        this.cdr.markForCheck();
      },
      (error) => {
        console.error('Error fetching profile', error);
      }
    );

  }

  fetchPosts() {

    this.http.get<any[]>('http://127.0.0.1:4201/getPosts/testuser').subscribe(
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
}
