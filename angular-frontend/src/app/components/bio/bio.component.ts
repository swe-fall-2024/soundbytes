import {Component, ElementRef, Input, Renderer2} from '@angular/core';
import {MatIconModule} from '@angular/material/icon';
import {MatChipsModule} from '@angular/material/chips';
import {MatCardModule} from '@angular/material/card';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-bio',
  template: `
  <mat-card class="profile-card" >
    <mat-card-header>
      <div mat-card-avatar class="example-header-image"></div>
      <mat-card-title>{{profile?.name}} {{profile?.username}}</mat-card-title>
      <mat-card-subtitle>{{profile?.currentFavType}}: {{profile?.currentFav}}</mat-card-subtitle>
      @if (me) {
      <mat-card-actions align="end">
            <button class="btn btn-outline-success search-button" routerLink="/edit-profile" type="submit">Edit</button>
      </mat-card-actions>
      }
      @else {
        @if (following) {
        <mat-card-actions align="end">
              <button class="btn btn-outline-success search-button" type="submit">Following</button>
        </mat-card-actions>
        }
        @else {
        <mat-card-actions align="end">
              <button class="btn btn-outline-success search-button" type="submit">Follow</button>
        </mat-card-actions>
        }
      }
    </mat-card-header>
    <mat-card-content>
      <p>
        
      </p>
      <mat-chip-set>
        <mat-chip>{{profile.genres[0]}}</mat-chip>
        <mat-chip>{{profile.genres[1]}}</mat-chip>
        <mat-chip>{{profile.genres[2]}}</mat-chip>
      </mat-chip-set>
      <br>
      <div fxLayout="row" fxLayoutAlign="start center">
          <mat-icon>music_note</mat-icon>
          <span>{{profile?.topSong}}</span>
          &nbsp;
          <mat-icon>person</mat-icon>
          <span>{{profile?.topArtist}}</span>
        </div>    
    </mat-card-content>
    </mat-card>
    <router-outlet></router-outlet>
    `,
  imports: [MatCardModule, MatIconModule, MatChipsModule, RouterModule],
  styleUrl: './bio.component.css'
})
export class BioComponent {
  @Input() profile: any; 
  @Input() me: boolean | undefined;
  @Input() following: boolean | undefined;
  @Input() pfp: string | undefined;

  constructor(private el: ElementRef, private renderer: Renderer2) {}

  ngAfterViewInit(): void {
    console.log(this.profile.pic);
    const divElement = this.el.nativeElement.querySelector('.example-header-image');
    this.renderer.setStyle(divElement, 'background-image', `url(${this.profile.pic})`);
  }
}
