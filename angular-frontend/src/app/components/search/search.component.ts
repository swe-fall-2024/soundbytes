import { Component } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { NgFor,NgIf } from '@angular/common';
import { FormControl, FormsModule } from '@angular/forms';
import {MatIconModule} from '@angular/material/icon';
import {MatButtonModule} from '@angular/material/button';
import {MatToolbarModule} from '@angular/material/toolbar';
import {AsyncPipe} from '@angular/common';
import {MatAutocompleteModule} from '@angular/material/autocomplete';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import { UserService } from '../../services/signup.component'; // Import the service
import { HttpClient, HttpClientModule } from '@angular/common/http';


@Component({
  selector: 'app-search',
  imports: [NgFor, NgIf,FormsModule, MatToolbarModule, MatButtonModule, MatIconModule],
  templateUrl: './search.component.html',
  styleUrl: './search.component.css'
})
export class SearchComponent {

  searchText = '';
  items = [
    { username: 'mary' },
    { username: 'matt' },
    { username: 'katie' },
    { username: 'gaby' }
  ];

  constructor(private http: HttpClient, private router: Router, private userService: UserService) {}

  get filteredItems() {
    //if user has no input -> placeholder search... + no results
    if(this.searchText == ''){
      return null
    }

    //get results from database based on search text inputed by user
    var results = this.items.filter(item => 
      item.username.toLowerCase().includes(this.searchText.toLowerCase())
    );

    if(results.length == 0){
      return null
    }

    return results
  }

  navigateToUser(username: any){
    console.log('in navigate to user')
    this.searchText = '';
    
    console.log(username)
    this.router.navigate(['/friend-profile', username]);
  }
}
