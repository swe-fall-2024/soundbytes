import {Component, Input} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatListModule} from '@angular/material/list';
import {MatDividerModule} from '@angular/material/divider';

@Component({
  selector: 'app-friends',
  template: `
  <mat-list-item>
        <div fxLayout="row" fxLayoutAlign="start center" class="icon-text">
          <img src="assets/download.jpg" class="icon-image">
          <span>{{friend.name}}</span>
        </div>
      </mat-list-item>
  <mat-divider></mat-divider>
  `,
  imports: [MatButtonModule, MatCardModule, MatListModule, MatDividerModule],
  styleUrl: './friends.component.css',
})
export class FriendsComponent {
  @Input() friend: any;
}
