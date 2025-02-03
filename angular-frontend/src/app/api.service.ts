import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

interface MessageResponse {
  message: string;
}

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private apiUrl = 'http://localhost:8080/'; // ✅ Ensure this matches your Go backend

  constructor(private http: HttpClient) {}

  getMessage(): Observable<MessageResponse> {
    return this.http.get<MessageResponse>(this.apiUrl);
  }
}
