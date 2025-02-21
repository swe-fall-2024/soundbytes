import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppComponent } from './app.component';
import { provideHttpClient } from '@angular/common/http';
import { provideHttpClientTesting } from '@angular/common/http/testing';
//import { ApiService } from './api.service'; // ✅ Import service

@NgModule({
  declarations: [AppComponent],
  imports: [BrowserModule], 
  providers: [provideHttpClient(),
              provideHttpClientTesting()
  ], // ✅ Use `provideHttpClient()`
  bootstrap: [AppComponent],
})
export class AppModule {}
