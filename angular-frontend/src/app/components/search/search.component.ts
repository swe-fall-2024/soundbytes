import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { NgFor, NgIf } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatAutocompleteModule } from '@angular/material/autocomplete';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { SearchService } from '../../search.service'; // Import the service
import { User } from '../../models/user.model'; // Import the User model

@Component({
  selector: 'app-search',
  imports: [
    NgFor,
    NgIf,
    FormsModule,
    MatToolbarModule,
    MatButtonModule,
    MatIconModule,
    MatAutocompleteModule,
    MatInputModule,
    MatFormFieldModule
  ],
  templateUrl: './search.component.html',
  styleUrls: ['./search.component.css']
})
export class SearchComponent {
  searchText: string = '';  // User input
  filteredItems: User[] = []; // Filtered users list
  allUsers: User[] = []; // Initially empty array

  constructor(private searchService: SearchService, private router: Router) {}

  // Method to filter users based on the search text
  filterUsers() {
    if (this.searchText === '') {
      this.filteredItems = [];
    } else {
      // Call the SearchService to get the filtered users
      this.searchService.searchUsers(this.searchText).subscribe(
        (response) => {
          this.filteredItems = response; // Assuming response is an array of User objects
        },
        (error) => {
          console.error('Error fetching users:', error);
          this.filteredItems = [];
        }
      );
    }
  }

  // Navigate to the user's profile page
  navigateToUser(userID: any) {
    console.log("Navigating to user:", userID);
    this.searchText = '';  // Clear the search text after navigation
    this.router.navigate(['/friend-profile', userID]);  // Use Angular Router for navigation
  }
}
