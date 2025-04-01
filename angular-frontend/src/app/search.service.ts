import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, catchError, throwError, tap } from 'rxjs';
import { User } from './models/user.model';  // Import the User model from the correct location

@Injectable({
  providedIn: 'root'
})
export class SearchService {
  private apiUrl = 'http://localhost:4201/searchUsers'; // Backend URL

  constructor(private http: HttpClient) {}

  // Method to search for users based on the query
  searchUsers(query: string): Observable<User[]> {
    return this.http.get<User[]>(`${this.apiUrl}?q=${query}`).pipe(
      tap((data) => {
        console.log('Data received in SearchService:', data); // Log to verify data
      }),
      catchError((error) => {
        console.error('Error fetching users in service:', error);
        return throwError(error);  // Ensure the error is thrown correctly
      })
    );
  }
}
