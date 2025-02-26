import { Component, ChangeDetectionStrategy } from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {FormsModule} from '@angular/forms';
import {NgFor, CommonModule } from '@angular/common';
import {MatSelectModule} from '@angular/material/select';
import { UserService } from '../../services/signup.component';
import { HttpClient } from '@angular/common/http';
@Component({
  selector: 'app-edit-profile',
  imports: [MatSelectModule, NgFor, CommonModule, MatCardModule, MatInputModule, MatFormFieldModule, FormsModule],
  templateUrl: './edit-profile.component.html',
  styleUrl: './edit-profile.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})

export class EditProfileComponent {


  constructor(private http: HttpClient, private userService: UserService) {}

  ngOnInit() {
    this.bio[0].username = this.userService.getUsername();
  }

  bio = [
    {
      name: '',
      username: '',
      currentFavType: '',
      currentFav: '',
      genres: {one:'',two:'',three:''},
      topSong: "",
      topArtist: "",
    }
  ];

  submitProfile() {

    const profileData = {
      username: this.bio[0].username,
      currentFavType: this.bio[0].currentFavType,
      currentFav: this.bio[0].currentFav,
      genres: this.bio[0].genres,
      topSong: this.bio[0].topSong,
      topArtist: this.bio[0].topArtist
    };

    this.http.put('http://127.0.0.1:4201/setUpProfile', profileData).subscribe({

      next: (response) => {
        console.log('Profile setup successful', response);
        alert('Profile updated successfully!');
      },
      error: (error) => {
        console.error('Profile setup failed', error);
      }

    });
  }


}


