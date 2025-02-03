import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http'; // ✅ Required

import { AppComponent } from './app.component';
import { ApiService } from './api.service'; // ✅ Import service

@NgModule({
  declarations: [AppComponent],
  imports: [BrowserModule, HttpClientModule], // ✅ Provide HttpClient
  providers: [ApiService], // ✅ Register service
  bootstrap: [AppComponent]
})
export class AppModule {}
