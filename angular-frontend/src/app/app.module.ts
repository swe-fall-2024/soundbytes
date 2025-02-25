import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppComponent } from './app.component';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
import { AppRoutingModule } from './app-rounting.module';
import { FeedComponent } from './views/feed/feed.component'
import { EditProfileComponent } from './views/edit-profile/edit-profile.component';
import { LoginComponent } from './views/login/login.component';
import { SignupComponent } from './views/signup/signup.component';
import { ProfileComponent } from './views/profile/profile.component';
import { PostCreationComponent } from './views/post-creation/post-creation.component';
import { provideRouter, RouterModule } from '@angular/router';
//import { ApiService } from './api.service'; // ✅ Import service

@NgModule({
  declarations: [
    AppComponent,
    FeedComponent,
    EditProfileComponent,
    LoginComponent,
    SignupComponent,
    ProfileComponent,
    PostCreationComponent
  ],
  imports: [BrowserModule, RouterModule ], 
  providers: [provideHttpClient(),
              provideHttpClientTesting(),
  ], // ✅ Use `provideHttpClient()`
  bootstrap: [AppComponent],
})
export class AppModule {}
