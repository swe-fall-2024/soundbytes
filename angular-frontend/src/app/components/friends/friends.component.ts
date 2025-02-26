import {Component, Input} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatListModule} from '@angular/material/list';
import {MatDividerModule} from '@angular/material/divider';
import { RouterLink } from '@angular/router';

@Component({
  selector: 'app-friends',
  template: `
  <mat-list-item>
        <div fxLayout="row" fxLayoutAlign="start center" class="icon-text">
          <ul class="navbar-nav mr-auto">
          <li class="nav-item"><a  class="nav-link" routerLink="/friend-profile" >{{friend.name}}</a></li> 
          </ul>
        </div>
      </mat-list-item>
  <mat-divider></mat-divider>
  `,
  imports: [MatButtonModule, MatCardModule, MatListModule, MatDividerModule, RouterLink],
  styleUrl: './friends.component.css',
})
export class FriendsComponent {
  @Input() friend: any;
}
