import {ChangeDetectionStrategy, Component, Input} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import { RouterModule } from '@angular/router';


@Component({
  selector: 'app-playlist-card',
    template: `
      <mat-card class="example-card" appearance="outlined">
        <mat-card-header>
          <mat-card-title>{{post?.title}}</mat-card-title>
          <mat-card-subtitle id="user" routerLink="/friend-profile/{{post?.user}}">{{post?.user}}</mat-card-subtitle>
          <div mat-card-avatar class="example-header-image"></div>
        </mat-card-header>
        <mat-card-content>
        <div>
          <mat-card appearance="outlined" class="song-player">
            <mat-card-content class="content">
              {{post?.content.playlist_title}}
              <a href={{post?.content.playlist_url}} target="_blank"  mat-button>
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
          <router-outlet></router-outlet>
    `,
    styleUrl: './playlist-card.component.css',
    imports: [MatCardModule, MatButtonModule, RouterModule],
    changeDetection: ChangeDetectionStrategy.OnPush,
  })

  export class PlaylistCardComponent {
    @Input() post: any;
  }

  /*
  embedded:
  <iframe style="border-radius:12px" src="https://open.spotify.com/embed/playlist/1yJb4XCnM4KfeO2UkMAYnp?utm_source=generator" width="100%" height="352" frameBorder="0" allowfullscreen="" allow="autoplay; clipboard-write; encrypted-media; fullscreen; picture-in-picture" loading="lazy"></iframe>
  non-embeded content:
  {{post.content.playlist_title}}
  <a href={{post.content.playlist_url}} target="_blank"  mat-button>
  <span class="material-icons">play_circle</span>
  </a>
  */