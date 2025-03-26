import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http'; 
import { FriendProfileComponent } from './friend-profile.component';

import { ActivatedRoute, convertToParamMap, RouteConfigLoadEnd } from '@angular/router';
import { async, Observable, of, Subject } from 'rxjs';

describe('FriendProfileComponent', () => {
  let component: FriendProfileComponent;
  let fixture: ComponentFixture<FriendProfileComponent>;
  const mockActivatedRoute = jasmine.createSpyObj({paramMap: new Observable()}, {params: { id: 1 }})
  

  beforeEach(async () => {
   
    await TestBed.configureTestingModule({
      imports: [FriendProfileComponent, HttpClientModule],
      providers: [
        {
          provide: ActivatedRoute,
          useValue: mockActivatedRoute
        }
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FriendProfileComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
