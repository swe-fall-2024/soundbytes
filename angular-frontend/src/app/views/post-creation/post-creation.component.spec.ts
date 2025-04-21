import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { PostCreationComponent } from './post-creation.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

describe('PostCreationComponent', () => {
  let component: PostCreationComponent;
  let fixture: ComponentFixture<PostCreationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PostCreationComponent, HttpClientModule, BrowserAnimationsModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PostCreationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
