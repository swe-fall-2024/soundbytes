import {ChangeDetectionStrategy, Component, OnInit} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatChipsModule} from '@angular/material/chips';
import {MatProgressBarModule} from '@angular/material/progress-bar';
import {MatIconModule} from '@angular/material/icon';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatListModule} from '@angular/material/list';
import {MatDividerModule} from '@angular/material/divider';
import {MatToolbarModule} from '@angular/material/toolbar';
import { BioComponent } from '../components/bio/bio.component';
import { NgIf, NgFor, CommonModule } from '@angular/common';
import { SongCardComponent } from '../components/song-card/song-card.component';
import { ReviewCardComponent } from '../components/review-card/review-card.component';
import { PlaylistCardComponent } from '../components/playlist-card/playlist-card.component';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-friend-profile',
  imports: [BioComponent, NgIf, NgFor, CommonModule, SongCardComponent, ReviewCardComponent, PlaylistCardComponent, MatToolbarModule, MatDividerModule, MatListModule, MatCardModule, MatButtonModule, MatChipsModule, MatProgressBarModule, MatIconModule, MatGridListModule],
  templateUrl: './friend-profile.component.html',
  styleUrl: './friend-profile.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class FriendProfileComponent implements OnInit {
  userId: string | null = '';

  constructor(private route: ActivatedRoute) {}

  ngOnInit() {
    // Access route parameter
    this.userId = this.route.snapshot.paramMap.get('id');
    console.log(this.userId)
  }
  
  posts = [
    {
      user: 'Cameron Santiago',
      profile_img: 'url',
      type: 'favorite-song',
      title: 'MY FAVORITE SONG',
      content:{
        song_title: 'Engagement Party',
        song_url: 'https://open.spotify.com/track/5PYPCxyWltRIyPkhSsnWIk',
        song_embed: "https://open.spotify.com/embed/track/6LxcPUqx6noURdA5qc4BAT?utm_source=generator",
        //song_embed
      }
    },
    {
      user: 'Cameron Santiago',
      profile_img: 'url',
      type: 'album-review',
      title: 'ALBUM REVIEW',
      content:{
        album_title: "Cats in the Cradle",
        review: "Harry Chapin’s iconic ballad Cats in the Cradle, released in 1974 on his Verities & Balderdash album, remains one of the most poignant reflections on fatherhood and the passage of time. Framed through deceptively simple lyrics and a memorable acoustic melody, the song captures the emotional distance that can grow between parents and children. Its timeless message about missed moments and generational echoes has resonated with listeners for decades, cementing its place as a powerful staple in American folk-rock storytelling.",
      }
    },
    {
      user: 'Shiba Inu',
      profile_img: 'url',
      type: 'playlist',
      title: 'MY NEW PLAYLIST',
      content:{
        playlist_title: "Rap playlist",
        playlist_url: "https://open.spotify.com/playlist/1yJb4XCnM4KfeO2UkMAYnp?si=945e25fe87034d38",
        playlist_embed: "https://open.spotify.com/embed/playlist/1yJb4XCnM4KfeO2UkMAYnp?utm_source=generator",
      }
    },
  ];
  profiles = [
    {
      name: 'Cameron Santiago',
      username: '@camsanti37',
      currentFavType: 'Current Favorite Artist',
      currentFav: 'Nirvana',
      genres: {one:'rap',two:'pop',three:'rock'},
      topSong: "Teenage Spirit",
      topArtist: "Harry Chaplin",
    }
  ];
  me = false;
  following = false;
}
