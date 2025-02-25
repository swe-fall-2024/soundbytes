import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface MessageResponse {
  message: string;
}

@Injectable({
  providedIn: 'root',
})
export class ApiService {
//  private readonly apiUrl = 'http://localhost:4201/';
  private readonly apiUrl = 'http://localhost:4201/api';

  constructor(private http: HttpClient) {}

  // Use Observable instead of Promise
  getMessage(): Observable<MessageResponse> {
    return this.http.get<MessageResponse>(this.apiUrl); // HttpClient directly returns an Observable
  }
}
