import { MatCardModule } from '@angular/material/card';
import { MatInputModule } from '@angular/material/input';
import { MatFormFieldModule } from '@angular/material/form-field';
import { ChangeDetectionStrategy, Component, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { FormControl, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatIconModule } from '@angular/material/icon';
import { merge } from 'rxjs';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Router, RouterLink, RouterOutlet } from '@angular/router';
import { UserService } from '../../services/signup.component'; // Import the service
import { NgIf } from '@angular/common';

@Component({
  selector: 'app-signup',
  standalone: true,
  imports: [MatCardModule, MatInputModule, MatFormFieldModule, FormsModule, ReactiveFormsModule, MatIconModule, HttpClientModule, NgIf, RouterLink,RouterOutlet],
  templateUrl: './signup.component.html',
  styleUrl: './signup.component.css',
  changeDetection: ChangeDetectionStrategy.OnPush,

})
export class SignupComponent {
  readonly email = new FormControl('', [Validators.required, Validators.email]);
  readonly password = new FormControl('', [Validators.required]);  // Ensure required validator for password

  errorMessage = signal('');

  constructor(private http: HttpClient, private router: Router, private userService: UserService) {
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
      this.errorMessage.set('You must enter a password');
    }
  }

  hide = signal(true);
  clickEvent(event: MouseEvent) {
    this.hide.set(!this.hide());
    event.stopPropagation();
  }

  // New method to handle signup
  signUp() {
    if (this.email.valid && this.password.valid) {
      const user = {
        username: this.email.value,
        password: this.password.value,
      };

      this.http.post('http://127.0.0.1:4201/register', user).subscribe({
        next: (response) => {
          if(user.username != null)
            localStorage.setItem('username', user.username);

          console.log('Registration successful', response);
          alert('Registration successful!');
          this.userService.setUsername(user.username || '');
          this.router.navigate(['/edit-profile']);
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
