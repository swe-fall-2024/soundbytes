import {ChangeDetectionStrategy, Component} from '@angular/core';
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
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { BioComponent } from '../../components/bio/bio.component';

@Component({
  selector: 'app-profile',
  imports: [BioComponent, NavbarComponent, FriendsComponent, NgIf, NgFor, CommonModule, SongCardComponent, ReviewCardComponent, PlaylistCardComponentComponent, MatToolbarModule, MatDividerModule, MatListModule, MatCardModule, MatButtonModule, MatChipsModule, MatProgressBarModule, MatIconModule, MatGridListModule],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ProfileComponent {
  posts = [
    {
      user: 'Shiba Inu',
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
      user: 'Shiba Inu',
      profile_img: 'url',
      type: 'album-review',
      title: 'ALBUM REVIEW',
      content:{
        album_title: "Short n' Sweet",
        review: "Sabrina Carpenter's latest album, Short n' Sweet, released on August 23, 2024, marks her sixth studio endeavor and showcases a refreshingly lighthearted and cheeky approach to pop music. The album has been lauded for its cleverness and effortless execution, setting a high bar for contemporary pop.",
      }
    },
    {
      user: 'Shiba Inu',
      profile_img: 'url',
      type: 'playlist',
      title: 'MY NEW PLAYLIST',
      content:{
        playlist_title: "Study playlist",
        playlist_url: "https://open.spotify.com/playlist/1yJb4XCnM4KfeO2UkMAYnp?si=945e25fe87034d38",
        playlist_embed: "https://open.spotify.com/embed/playlist/1yJb4XCnM4KfeO2UkMAYnp?utm_source=generator",
      }
    },
  ];
  friends = [
    {
      name: 'Katie'
    },
    {
      name: 'Mary'
    }
  ];
  profiles = [
    {
      name: 'Shiba Inu',
      currentFavType: 'Current Favorite Artist',
      currentFav: 'COIN',
      genres: {one:'indie',two:'pop',three:'hyperpop'},
      topSong: "stupid horse",
      topArtist: "100 gecs",
    }
  ];
}
