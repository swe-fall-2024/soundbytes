//   // src/app/app.component.ts
//   import { Component, OnInit } from '@angular/core';
//   import { ApiService } from './api.service'; // Import the ApiService

//   @Component({
//     selector: 'app-root',
//     templateUrl: './app.component.html',
//     styleUrls: ['./app.component.css'],
//   })
//   export class AppComponent implements OnInit {
//     title = 'angular-frontend';
//     message: string = '';

//     constructor(private apiService: ApiService) {} // Inject ApiService

//     ngOnInit() {
//       this.apiService.getMessage().subscribe({
//         next: (data) => {
//           this.message = data.message; // Set the message from the backend
//         },
//         error: (err) => {
//           console.error('Error fetching data:', err);
//         },
//       });
//     }
//   }




import { Component } from '@angular/core';
import {RouterModule } from '@angular/router';
import { NavbarComponent } from './components/navbar/navbar.component';
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
  imports: [RouterModule, NavbarComponent],
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


// import { Component } from '@angular/core';

// @Component({
//   selector: 'app-root',
//   template: `
//     <div>
//       <label for="textbox">Enter something:</label>
//       <input id="textbox" type="text" />
//     </div>
//   `,
//   styleUrls: ['./app.component.css']
// })
// export class AppComponent {
//   title = 'textbox-app';
// }




// import { Component, OnInit } from '@angular/core';
// import { ApiService } from './api.service';

// @Component({
//   selector: 'app-root',
//   template: `<h1>{{ message }}</h1>`,
// })
// export class AppComponent implements OnInit {
//   message: string = '';

//   constructor(private apiService: ApiService) {}

//   ngOnInit() {
//     this.apiService.getMessage().subscribe({
//       next: (data) => {
//         this.message = data.message; // ✅ Use the returned message
//       },
//       error: (err) => {
//         console.error('Error fetching data:', err);
//       }
//     });
//   }
// }
