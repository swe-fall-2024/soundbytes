import { ChangeDetectorRef, Component, OnInit } from '@angular/core';
import { SongCardComponent } from '../../components/song-card/song-card.component';
import { ReviewCardComponent } from '../../components/review-card/review-card.component';
import { PlaylistCardComponent } from '../../components/playlist-card/playlist-card.component';
import { NgIf, NgFor, CommonModule } from '@angular/common';
import { NavbarComponent } from '../../components/navbar/navbar.component';
import { ProfileService } from '../../profile.service';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-feed',
  imports: [NgIf, NgFor, NavbarComponent, CommonModule, SongCardComponent, ReviewCardComponent, PlaylistCardComponent],
  templateUrl: './feed.component.html',
  styleUrl: './feed.component.css'
})
export class FeedComponent implements OnInit {

  constructor(private profileService: ProfileService, private http: HttpClient, private cdr: ChangeDetectorRef) {}
  
  posts: any[] = [];

  ngOnInit(): void {
    this.fetchFeed()
  }


 
  fetchFeed() {

    this.http.get<any[]>(`http://127.0.0.1:4201/getFeed/${localStorage.getItem('currentUserEmail')}`).subscribe(
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

/* we can embed tracks and playlists from spotify:
<iframe style="border-radius:12px" src="https://open.spotify.com/embed/track/6LxcPUqx6noURdA5qc4BAT?utm_source=generator" width="100%" height="352" frameBorder="0" allowfullscreen="" allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture" loading="lazy"></iframe>
<iframe style="border-radius:12px" src="https://open.spotify.com/embed/playlist/1yJb4XCnM4KfeO2UkMAYnp?utm_source=generator" width="100%" height="352" frameBorder="0" allowfullscreen="" allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture" loading="lazy"></iframe>
*/
