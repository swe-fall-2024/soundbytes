import { Component, ChangeDetectionStrategy, Input } from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-review-card',
  template: `
    <mat-card class="example-card" appearance="outlined">
      <mat-card-header>
        <mat-card-title>{{post?.title}}</mat-card-title>
          <mat-card-subtitle id="user" routerLink="/friend-profile">{{post?.user}}</mat-card-subtitle>
        <div mat-card-avatar class="example-header-image"></div>
      </mat-card-header>
      <mat-card-content>
      <div>
        <mat-card appearance="outlined" class="song-player">
          <mat-card-content class="content">
            <h5>{{post?.content.album_title}}</h5>
            
            <p>{{post?.content.review}}</p>
            
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
  styleUrl: './review-card.component.css',
  imports: [MatCardModule, MatButtonModule, RouterModule],
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ReviewCardComponent {
  constructor() {
    console.log('myCustomComponent');
  }
  @Input() post: any;
}
