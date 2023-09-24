import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  private apiURL = 'http://localhost:8080/consulta1'
  constructor(private http: HttpClient) { }
  fetchData(URL: string): Observable<any> {
    // console.log(this.http.get(URL));
    return this.http.get(URL);
    
  }
  
}
