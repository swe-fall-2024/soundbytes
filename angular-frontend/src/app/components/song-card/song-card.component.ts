import {ChangeDetectionStrategy, Component, ElementRef, Input, Renderer2} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import { RouterModule } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { Router } from '@angular/router';

/**
 * @title Card with multiple sections
 */
@Component({
  selector: 'app-song-card',
  standalone: true,
  template: `
    <mat-card class="example-card" appearance="outlined">
      <mat-card-header>
        <mat-card-title>{{post?.title}}</mat-card-title>
        <mat-card-subtitle id="user">{{post?.user}}</mat-card-subtitle>
        <div mat-card-avatar class="example-header-image"></div>
      </mat-card-header>
      <mat-card-content>
        <div>
          <mat-card appearance="outlined" class="song-player">
            <mat-card-content class="content">
              {{post?.content.song_title}}
              <a href={{post?.content.song_url}} target="_blank" mat-button>
                <span class="material-icons">play_circle</span>
              </a>
            </mat-card-content>
          </mat-card>
        </div>
      </mat-card-content>
      <mat-card-actions>
        <button mat-button (click)="onLike()">LIKE</button>
        <span>  Like Count: {{post?.like_count || 0}}</span>
      </mat-card-actions>
    </mat-card>
    <router-outlet></router-outlet>
  `,
  styleUrl: './song-card.component.css',
  imports: [MatCardModule, MatButtonModule, RouterModule],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SongCardComponent {
  @Input() post: any;
  @Input() profile: any; 
  @Input() me: boolean | undefined;
  @Input() following: boolean | undefined;
  @Input() pfp: string | undefined;

  constructor(private http: HttpClient, private router: Router, private el: ElementRef, private renderer: Renderer2) {}
 
  ngAfterViewInit(): void {
    console.log(this.post.Profile_Image);
    const divElement = this.el.nativeElement.querySelector('.example-header-image');
    this.renderer.setStyle(divElement, 'background-image', `url(${this.post.Profile_Image})`);
  }

  onLike(): void {
    // Call a service to update the like count on the backend
    console.log('Like button clicked for post:', this.post);
    //alert(`I liked a Post! Testing! ${JSON.stringify(this.post.post_id)}`);
    // Example API call to update like count
   // alert(`Checking Request ${this.post.post_id}`);
    this.http.get<any[]>(`http://127.0.0.1:4201/likePost/${JSON.stringify(this.post.post_id)}`).subscribe(response => {
       //alert('Post liked successfully!');
       //alert(`Response: ${JSON.stringify(response)}`);
       if (this.post) {
         this.post.like_count = (this.post.like_count || 0) + 1;
       }
    }, 
    error => {
       console.error('Error updating like:', error);
    });
  }
}

/*
embedded:
<iframe style="border-radius:12px" src="https://open.spotify.com/embed/track/6LxcPUqx6noURdA5qc4BAT?utm_source=generator" width="100%" height="352" frameBorder="0" allowfullscreen="" allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture" loading="lazy"></iframe>
non-embedded content:
{{post.content.song_title}}
<a href={{post.content.song_url}} target="_blank" mat-button>
<span class="material-icons">play_circle</span>
</a>
*/