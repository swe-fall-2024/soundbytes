import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { provideHttpClient } from '@angular/common/http';

//import { ApiService } from './api.service'; // ✅ Import service

@NgModule({
  imports: [BrowserModule], 
  providers: [provideHttpClient()], // ✅ Use `provideHttpClient()`
})
export class AppModule {}
