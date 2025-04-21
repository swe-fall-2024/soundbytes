import { Component } from '@angular/core';
import {RouterModule } from '@angular/router';
import {
  trigger,
  state,
  style,
  animate,
  transition,
  // ...
} from '@angular/animations';


@Component({
  selector: 'app-root',
  imports: [RouterModule],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  animations: [
    trigger('fadeInOut', [  // Correctly defining the trigger
      state('void', style({ opacity: 0 })),  // Initial state
      transition(':enter', [animate('500ms ease-in', style({ opacity: 1 }))]), // Fade in
      transition(':leave', [animate('500ms ease-out', style({ opacity: 0 }))]) // Fade out
    ])
  ]
})
export class AppComponent {
  title = 'soundbytes';
}
