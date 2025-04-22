import { Component, ChangeDetectionStrategy, Input, ElementRef, Renderer2 } from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-review-card',
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
  @Input() post: any;
  @Input() profile: any; 
  @Input() me: boolean | undefined;
  @Input() following: boolean | undefined;
  @Input() pfp: string | undefined;

  constructor(private el: ElementRef, private renderer: Renderer2) {}

  ngAfterViewInit(): void {
    console.log(this.post.Profile_Image);
    const divElement = this.el.nativeElement.querySelector('.example-header-image');
    this.renderer.setStyle(divElement, 'background-image', `url(${this.post.Profile_Image})`);
  }
}
