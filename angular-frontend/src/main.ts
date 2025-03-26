import { bootstrapApplication } from '@angular/platform-browser';
import { AppComponent } from './app/app.component';
import { provideHttpClient } from '@angular/common/http';
import { provideAnimations } from '@angular/platform-browser/animations';
import { appConfig } from './app/app.config';
import { provideNoopAnimations } from '@angular/platform-browser/animations';


bootstrapApplication(AppComponent, {
  providers: [...appConfig.providers,provideHttpClient(), provideAnimations(), provideNoopAnimations()]
})
  .catch(err => console.error('Bootstrap error:', err));
