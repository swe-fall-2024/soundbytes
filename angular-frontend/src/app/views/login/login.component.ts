import {MatCardModule} from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {ChangeDetectionStrategy, Component, signal} from '@angular/core';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {FormControl, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatIconModule} from '@angular/material/icon';
import {merge} from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login',
  imports: [MatCardModule, MatInputModule, MatFormFieldModule, FormsModule, ReactiveFormsModule,MatIconModule],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginComponent {
  readonly email = new FormControl('', [Validators.required, Validators.email]);
  readonly password = new FormControl('', );

  errorMessage = signal('');

  constructor(private http: HttpClient) {
    merge(this.email.statusChanges, this.email.valueChanges)
      .pipe(takeUntilDestroyed())
      .subscribe(() => this.updateErrorMessage());

    merge(this.password.statusChanges, this.password.valueChanges)
      .pipe(takeUntilDestroyed())
      .subscribe(() => this.updateErrorMessage());
  }

  updateErrorMessage() {
    if (this.email.hasError('required')) {
      this.errorMessage.set('You must enter a value');
    } else if (this.email.hasError('email')) {
      this.errorMessage.set('Not a valid email');
    } else {
      this.errorMessage.set('');
    }

    if (this.password.hasError('required')) {
      this.errorMessage.set('You must enter a value');
    }
  }

  hide = signal(true);
  clickEvent(event: MouseEvent) {
    this.hide.set(!this.hide());
    event.stopPropagation();
  }


  // New method to handle signup
  login() {
    if (this.email.valid && this.password.valid) {
      const user = {
        username: this.email.value,
        password: this.password.value,
      };

      this.http.post('http://127.0.0.1:4201/login', user).subscribe({
        next: (response) => {
          console.log('Login successful', response);
          alert('Login successful!');
        },
        error: (error) => {
          console.error('Registration failed', error);
          this.errorMessage.set('Registration failed. Please try again.');
        }
      });
    } else {
      this.updateErrorMessage();
    }
  }
}
