import { Component } from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {FormsModule} from '@angular/forms';
import {NgFor, CommonModule } from '@angular/common';
import {MatSelectModule} from '@angular/material/select';

@Component({
  selector: 'app-edit-profile',
  imports: [MatSelectModule, NgFor, CommonModule, MatCardModule, MatInputModule, MatFormFieldModule, FormsModule],
  templateUrl: './edit-profile.component.html',
  styleUrl: './edit-profile.component.css'
})
export class EditProfileComponent {
  bio = [
    {
      name: 'Shiba Inu',
      username: '@shiba',
      currentFavType: 'Current Favorite Artist',
      currentFav: 'COIN',
      genres: {one:'indie',two:'pop',three:'hyperpop'},
      topSong: "stupid horse",
      topArtist: "100 gecs",
    }
  ];
}
