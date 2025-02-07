import {ChangeDetectionStrategy, Component, inject} from '@angular/core';
import {MatButtonModule} from '@angular/material/button';
import {MatCardModule} from '@angular/material/card';
import {MatChipsModule} from '@angular/material/chips';
import {MatProgressBarModule} from '@angular/material/progress-bar';
import {MatIconModule} from '@angular/material/icon';
import {MatGridListModule} from '@angular/material/grid-list';


@Component({
  selector: 'app-profile',
  imports: [MatCardModule, MatButtonModule, MatChipsModule, MatProgressBarModule, MatIconModule, MatGridListModule],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class ProfileComponent {
  
}

// import {ChangeDetectionStrategy, Component} from '@angular/core';
// import {MatButtonModule} from '@angular/material/button';
// import {MatCardModule} from '@angular/material/card';
// //import {MatChipsModule} from '@angular/material/chips';

// /**
//  * @title Card with multiple sections
//  */
// @Component({
//   selector: 'profile',
//   templateUrl: 'profile.html',
//   styleUrl: 'profile.css',
//   imports: [MatCardModule, MatButtonModule],
//   changeDetection: ChangeDetectionStrategy.OnPush,
// })
// export class profile {}
