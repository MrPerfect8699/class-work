import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';

import { AppRoutingModule } from './app-routing.module';
import { LoginComponent } from './components/login/login.component';
import { DashboardComponent } from './components/dashboard/dashboard.component';
import { HomeworkFormComponent } from './components/homework-form/homework-form.component';
import { HomeworkListComponent } from './components/homework-list/homework-list.component';
import { AppComponent } from './app.component';

@NgModule({
  declarations: [
    AppComponent, LoginComponent, DashboardComponent, HomeworkFormComponent, HomeworkListComponent
  ],
  imports: [
    BrowserModule, AppRoutingModule, HttpClientModule, FormsModule, BrowserAnimationsModule,
    MatInputModule, MatButtonModule, MatCardModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
