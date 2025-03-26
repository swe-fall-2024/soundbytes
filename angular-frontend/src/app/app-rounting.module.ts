import { NgModule } from '@angular/core'; 
import { Routes, RouterModule } from '@angular/router'; 
import { FeedComponent } from './views/feed/feed.component'
import { EditProfileComponent } from './views/edit-profile/edit-profile.component';
import { LoginComponent } from './views/login/login.component';
import { SignupComponent } from './views/signup/signup.component';
import { ProfileComponent } from './views/profile/profile.component';
import { PostCreationComponent } from './views/post-creation/post-creation.component';
import { BrowserModule } from '@angular/platform-browser';
import { AppComponent } from './app.component';
import { LandingPageComponent } from './views/landing-page/landing-page.component';

const routes: Routes = [ 
  { path: '', component: LandingPageComponent}, 
  { path: 'login', component: LoginComponent }, 
  { path: 'signup', component: SignupComponent }, 
  { path: 'feed', component: FeedComponent }, 
  { path: 'profile', component: ProfileComponent }, 
  { path: 'edit-profile', component: EditProfileComponent }, 
  { path: 'post', component: PostCreationComponent }, 
]; 

@NgModule({ 
imports: [BrowserModule, RouterModule.forRoot(routes)], 
exports: [RouterModule], 
providers: [] 
}) 
export class AppRoutingModule { } 
