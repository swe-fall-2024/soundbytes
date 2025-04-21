import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { LoginComponent } from './login.component';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('LoginComponent', () => {
  let component: LoginComponent;
  let fixture: ComponentFixture<LoginComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LoginComponent, HttpClientModule, BrowserAnimationsModule],
      providers: [
        {
          provide: ActivatedRoute,
          useValue: { params: of({ id: 123 }) }
        }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(LoginComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create login', () => {
    expect(component).toBeTruthy();
  });
});
