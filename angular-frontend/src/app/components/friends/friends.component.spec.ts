import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { FriendsComponent } from './friends.component';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';

describe('FriendsComponent', () => {
  let component: FriendsComponent;
  let fixture: ComponentFixture<FriendsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FriendsComponent, HttpClientModule],
      providers: [
        {
          provide: ActivatedRoute,
          useValue: { params: of({ id: 123 }) }
        }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FriendsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
