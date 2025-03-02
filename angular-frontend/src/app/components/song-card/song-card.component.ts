import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';

/**
 * @title Card with multiple sections
 */
@Component({
  selector: 'app-song-card',
  //templateUrl: './card.component.html',
  template: `
    <mat-card class="example-card" appearance="outlined">
      <mat-card-header>
        <mat-card-title>{{post?.title}}</mat-card-title>

          <mat-card-subtitle>{{post?.user}}</mat-card-subtitle>
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
        <button mat-button>LIKE</button>
        <button mat-button>SHARE</button>
      </mat-card-actions>
    </mat-card>

  `,
  styleUrl: './song-card.component.css',
  imports: [MatCardModule, MatButtonModule],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SongCardComponent {
  constructor() {
    console.log('myCustomComponent');
  }
  @Input() post: any;
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