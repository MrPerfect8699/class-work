import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { AuthService } from '../../auth.service';

@Component({
  templateUrl: './login.component.html'
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
