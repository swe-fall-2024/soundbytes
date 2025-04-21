import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { SignupComponent } from './signup.component';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('SignupComponent', () => {
  let component: SignupComponent;
  let fixture: ComponentFixture<SignupComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SignupComponent, HttpClientModule, BrowserAnimationsModule ],
      providers: [
        {
          provide: ActivatedRoute,
          useValue: { params: of({ id: 123 }) }
        }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SignupComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create signup component', () => {
    expect(component).toBeTruthy();
  });
});
