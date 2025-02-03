import { bootstrapApplication } from '@angular/platform-browser';
import { AppComponent } from './app/app.component';
import { provideHttpClient, HttpClient } from '@angular/common/http';
import { inject } from '@angular/core';

bootstrapApplication(AppComponent, {
  providers: [provideHttpClient()]
}).then(appRef => {
  const http = appRef.injector.get(HttpClient); // ✅ Inject HttpClient properly

  // Fetch data from Go backend
  http.get<{ message: string }>('http://localhost:8080/')
    .subscribe({
      next: (data) => {
        console.log('Response from Go:', data);
        document.body.innerHTML = `<h1>${data.message}</h1>`; // ✅ Render the response
      },
      error: (err) => console.error('Error fetching data:', err),
    });
}).catch(err => console.error('Bootstrap error:', err));
