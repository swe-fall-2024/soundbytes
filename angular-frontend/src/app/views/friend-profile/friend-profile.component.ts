import {ChangeDetectionStrategy, Component} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatChipsModule} from '@angular/material/chips';
import {MatIconModule} from '@angular/material/icon';
import {MatGridListModule} from '@angular/material/grid-list';
import {MatListModule} from '@angular/material/list';
import {MatDividerModule} from '@angular/material/divider';
import {MatToolbarModule} from '@angular/material/toolbar';

@Component({
  selector: 'app-friend-profile',
  imports: [MatToolbarModule, MatDividerModule, MatListModule, MatCardModule, MatButtonModule, MatChipsModule, MatIconModule, MatGridListModule],
  templateUrl: './friend-profile.component.html',
  styleUrl: './friend-profile.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class FriendProfileComponent {

}
