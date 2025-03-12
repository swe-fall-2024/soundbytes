import { Routes } from '@angular/router';
import { FeedComponent } from './views/feed/feed.component'
import { EditProfileComponent } from './views/edit-profile/edit-profile.component';
import { LoginComponent } from './views/login/login.component';
import { SignupComponent } from './views/signup/signup.component';
import { ProfileComponent } from './views/profile/profile.component';
import { PostCreationComponent } from './views/post-creation/post-creation.component';
import { LandingPageComponent } from './views/landing-page/landing-page.component';
import { FriendProfileComponent } from './friend-profile/friend-profile.component';
import { AuthGuard } from './services/guard.component';

export const routes: Routes = [
  //{ path: '', component: LandingPageComponent}, 
  { path: 'login', component: LoginComponent}, 
  { path: 'signup', component: SignupComponent}, 
  { path: '', component: FeedComponent}, 
  { path: 'profile', component: ProfileComponent }, 
<<<<<<< HEAD
  { path: 'edit-profile', component: EditProfileComponent}, 
  { path: 'post', component: PostCreationComponent},
  { path: 'friend-profile', component: FriendProfileComponent},
  { path: '**', redirectTo: '', pathMatch: 'full' }
=======
  { path: 'edit-profile', component: EditProfileComponent }, 
  { path: 'post', component: PostCreationComponent },
  //{ path: '', redirectTo: '', pathMatch: 'full' }
>>>>>>> mary-frontend
];

//if logged out -> redirect to login for base path, prevent access to other paths like feed
//if logged in -> redirect to feed on base path, redirect on navigate to login/signup
