import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { ReviewCardComponent } from './review-card.component';

describe('ReviewCardComponent', () => {
  let component: ReviewCardComponent;
  let fixture: ComponentFixture<ReviewCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ReviewCardComponent, HttpClientModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ReviewCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    //set data to test
    component.post = 
    {
      user: 'Anna',
      profile_img: 'url',
      type: 'album-review',
      title: 'ALBUM REVIEW',
      content:{
        album_title: "Short n' Sweet",
        review: "Sabrina Carpenter's latest album, Short n' Sweet, released on August 23, 2024, marks her sixth studio endeavor and showcases a refreshingly lighthearted and cheeky approach to pop music. The album has been lauded for its cleverness and effortless execution, setting a high bar for contemporary pop.",
      }
    };
    
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });
});
