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

@Component({
  selector: 'app-post-creation',
  imports: [MatCardModule, MatInputModule, MatFormFieldModule, FormsModule, ReactiveFormsModule,MatIconModule, MatSelect, MatOption, NgIf],
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
    console.log(this.selectedType.value);
  }
}
