import { Component } from '@angular/core';
import { RouterLink, RouterOutlet } from '@angular/router';
import { SearchComponent } from '../search/search.component';

@Component({
  selector: 'app-navbar',
  imports: [RouterLink, SearchComponent],
  templateUrl: './navbar.component.html',
  styleUrl: './navbar.component.css'
})
export class NavbarComponent {
  out = true;
  in = false;

  updateNavbar(){
    if(String(localStorage.getItem('currentUserEmail')) != null) {
      this.out = false;
      this.in = true;
    }
    else {
      this.out = true;
      this.in = false;
    }
  }

}
