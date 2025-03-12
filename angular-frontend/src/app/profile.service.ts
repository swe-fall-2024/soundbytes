import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { catchError, Observable, tap, throwError } from 'rxjs';
import { Profile } from './models/profile.model';

@Injectable({
  providedIn: 'root'
})
export class ProfileService {
  private apiUrl = 'http://localhost:4201/profile'; // Backend URL

  constructor(private http: HttpClient) {}

  // Use Profile type instead of any for better type safety
  getUserProfile(userID: string): Observable<Profile> {
    // Ensure the URL string is correct
    return this.http.get<Profile>(`${this.apiUrl}?userId=${userID}`).pipe(
      tap((data) => {
        console.log('Data received in ProfileService:', data); // Log to verify data
      }),
      catchError((error) => {
        console.error('Error fetching data in service:', error);
        return throwError(error);  // Ensure the error is thrown correctly
      })
    );
  }

  updateUserProfile(userID: string, updatedProfile: Profile): Observable<Profile> {
    return this.http.put<Profile>(`${this.apiUrl}?userId=${userID}`, updatedProfile).pipe(
      tap((data) => console.log('Profile updated successfully:', data)),
      catchError((error) => {
        console.error('Error updating profile:', error);
        return throwError(error);
      })
    );
  }
  
}
