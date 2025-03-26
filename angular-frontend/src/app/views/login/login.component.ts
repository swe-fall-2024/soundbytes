import {MatCardModule} from '@angular/material/card';
import {MatInputModule} from '@angular/material/input';
import {MatFormFieldModule} from '@angular/material/form-field';
import {ChangeDetectionStrategy, Component, signal} from '@angular/core';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {FormControl, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {MatIconModule} from '@angular/material/icon';
import {merge} from 'rxjs';
import { Router, RouterLink, RouterOutlet } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { NgIf } from '@angular/common';
import {
  trigger,
  state,
  style,
  animate,
  transition,
  // ...
} from '@angular/animations';

@Component({
  selector: 'app-login',
  imports: [MatCardModule, MatInputModule, MatFormFieldModule, FormsModule, ReactiveFormsModule,MatIconModule, NgIf, RouterLink, RouterOutlet],
  templateUrl: './login.component.html',
  styleUrl: './login.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class LoginComponent {
  readonly email = new FormControl('', [Validators.required, Validators.email]);
  readonly password = new FormControl('', );

  errorMessage = signal('');
  isLoggedIn = '';

  constructor(private http: HttpClient, private router: Router) {
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
    console.log("on submit i guess")
    this.hide.set(!this.hide());
    event.stopPropagation();
  }


  // New method to handle login
  login() {
    if (this.email.valid && this.password.valid) {
      const user = {
        username: this.email.value,
        password: this.password.value,
      };

      this.http.post('http://127.0.0.1:4201/login', user).subscribe({
        next: (response) => {
          if(user.username != null)
            localStorage.setItem('currentUserEmail', user.username);

          console.log('Login successful', response);
          alert('Login successful!');
          this.isLoggedIn = 'Login successful!';
          this.router.navigate(['/profile']);
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
