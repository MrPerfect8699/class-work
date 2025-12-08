import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../core/auth.service';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrl: './login.component.scss',
  standalone: true,
  imports: [CommonModule, FormsModule, MatCardModule, MatFormFieldModule, MatInputModule, MatButtonModule],
})
export class LoginComponent {
  email = '';
  password = '';
  name = '';
  isRegister = false;
  constructor(private auth: AuthService, private router: Router) {}

  toggle(){ this.isRegister = !this.isRegister; }

  submit(){
    if (this.isRegister){
      this.auth.register(this.name, this.email, this.password).subscribe(()=> {
        alert('registered — now login');
        this.isRegister = false;
      }, e => alert('err'));
    } else {
      this.auth.login(this.email, this.password).subscribe(()=> {
        this.router.navigate(['/dashboard']);
      }, e => alert('login failed'));
    }
  }
}
