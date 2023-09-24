import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { BotonesComponent } from './botones/botones.component';
import { ApiService } from './api.service';
import { HttpClientModule } from '@angular/common/http';
import { MatCardModule } from '@angular/material/card';


@NgModule({
  declarations: [
    AppComponent,
    BotonesComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    MatCardModule,
    
  ],
  providers: [ApiService],
  bootstrap: [AppComponent, BotonesComponent]
})
export class AppModule { }
