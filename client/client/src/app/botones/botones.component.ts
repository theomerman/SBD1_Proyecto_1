import { Component } from '@angular/core';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-botones',
  templateUrl: './botones.component.html',
  styleUrls: ['./botones.component.css']
})
export class BotonesComponent {
  fetchedData: any;


  constructor(private apiService: ApiService) {}

  fetchData(URL: string) {
    this.apiService.fetchData("http://localhost:8080/" + URL).subscribe((data) => {
      this.fetchedData = data;
      // this.fetchedData = '```json\n' + JSON.stringify(data, null, 2) + '\n```';

      console.log(this.fetchedData);
    });
  }
}
