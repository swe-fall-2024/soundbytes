import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

export interface MessageResponse {
  message: string;
}

@Injectable({
  providedIn: 'root' // ✅ Global service
})
export class ApiService {
  private readonly apiUrl = 'http://localhost:4201/'; // ✅ Readonly to prevent modifications
  private http = inject(HttpClient); // ✅ Use `inject` instead of constructor DI

  getMessage(): Observable<MessageResponse> {
    return this.http.get<MessageResponse>(this.apiUrl);
  }
}
