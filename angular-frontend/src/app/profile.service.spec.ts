import { TestBed } from '@angular/core/testing';

import { ProfileService } from './profile.service';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';
import { HttpClientModule } from '@angular/common/http';

describe('ProfileService', () => {
  let service: ProfileService;

  beforeEach(() => {
    TestBed.configureTestingModule(
      {
        imports: [HttpClientModule ],
        providers: [
              {
                provide: ActivatedRoute,
                useValue: { params: of({ id: 123 }) }
              }
            ]}
    );
    service = TestBed.inject(ProfileService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
