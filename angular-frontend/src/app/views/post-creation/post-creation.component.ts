import { Component } from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {ChangeDetectionStrategy, signal} from '@angular/core';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {FormControl,FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatIconModule} from '@angular/material/icon';
import {merge} from 'rxjs';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { MatOption, MatSelect } from '@angular/material/select';
import { NgIf } from '@angular/common';
import { NavbarComponent } from '../../components/navbar/navbar.component';

@Component({
  selector: 'app-post-creation',
  imports: [MatCardModule, MatInputModule, NavbarComponent, MatFormFieldModule, FormsModule, ReactiveFormsModule,MatIconModule, MatSelect, MatOption, NgIf],
  templateUrl: './post-creation.component.html',
  styleUrl: './post-creation.component.css'
})
export class PostCreationComponent {
  /*myForm = new FormGroup({
    selectedType: new FormControl(''),
    review: new FormControl(''),
    reviewTitle: new FormControl(''),
    songTitle: new FormControl(''),
    songLink: new FormControl(''),
    playlistTitle: new FormControl(''),
    playlistLink: new FormControl(''),
  });*/
  
  readonly selectedType = new FormControl('');
  readonly review = new FormControl('');
  readonly reviewTitle = new FormControl('');
  readonly songTitle = new FormControl('');
  readonly songLink = new FormControl('');
  readonly playlistTitle = new FormControl('');
  readonly playlistLink = new FormControl('');

  

  errorMessage = signal('');

  constructor(private http: HttpClient, private router: Router) {

    merge(this.selectedType.statusChanges, this.selectedType.valueChanges)
      .pipe(takeUntilDestroyed())
      .subscribe(() => this.updateErrorMessage());

  }

  updateErrorMessage() {

    if (this.selectedType.hasError('required')) {
      this.errorMessage.set('You must enter a value');
    } 

  }

  hide = signal(true);
  clickEvent(event: MouseEvent) {
    this.hide.set(!this.hide());
    event.stopPropagation();
  }

  
  submit(){

    let content: any = {};

    // Check the selected type and populate the content object accordingly
    if (this.selectedType.value === 'favorite-song') {
      content = {
        song_title: this.songTitle.value || '',
        song_url: this.songLink.value || '',
        song_embed: 'sample',
      };
    } 
    
    else if (this.selectedType.value === 'album-review') {
      content = {
        album_title: this.reviewTitle.value || '',
        review: this.review.value || '',
      };
    } 
    
    else if (this.selectedType.value === 'playlist') {
      content = {
        playlist_title: this.playlistTitle.value || '',
        playlist_url: this.playlistLink.value || '',
        playlist_embed: '',
      };
    }

    alert('New Post!');
    alert(`THIS IS A SIGN ${this.selectedType.value}`);
    alert(`Song Title: ${this.songTitle.value}`);
    alert(`Song Link: ${this.songLink.value}`);

    const post = {
      post_id: Math.floor(Math.random() * 200001),
      user: String(localStorage.getItem('currentUserEmail')), 
      profile_img: "",
      type: (this.selectedType.value) || '',
      title: (this.songTitle.value) || '',
      content: {
        song_title: this.songTitle.value || '',
        song_url: this.songLink.value || '',
        song_embed: '',
        album_title: this.reviewTitle.value || '',
        review: this.review.value || '',
        playlist_title: this.playlistTitle.value || '',
        playlist_url: this.playlistLink.value || '',
        playlist_embed: '',
      },
      like_count: 0
    };


    this.http.post('http://127.0.0.1:4201/addPost', post).subscribe({

      next: (response) => {
        console.log('Posts created successfully', response);
        alert('New Post!');
        this.router.navigate([`/profile`]);

      },
      error: (error) => {
        alert(`Error: ${JSON.stringify(error)}`);
        console.error('Registration failed', error);
        this.errorMessage.set('Registration failed. Please try again.');
      }
    });
  }
}
