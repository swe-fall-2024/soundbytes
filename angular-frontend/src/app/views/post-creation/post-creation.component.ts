import { Component } from '@angular/core';
import {MatCardModule} from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {ChangeDetectionStrategy, signal} from '@angular/core';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {FormControl, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatIconModule} from '@angular/material/icon';
import {merge} from 'rxjs';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { MatOption, MatSelect } from '@angular/material/select';

@Component({
  selector: 'app-post-creation',
  imports: [MatCardModule, MatInputModule, MatFormFieldModule, FormsModule, ReactiveFormsModule,MatIconModule, MatSelect, MatOption],
  templateUrl: './post-creation.component.html',
  styleUrl: './post-creation.component.css'
})
export class PostCreationComponent {
  readonly selectedType = new FormControl('');
  

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
