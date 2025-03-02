import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { SongCardComponent } from './song-card.component';

describe('SongCardComponent', () => {
  let component: SongCardComponent;
  let fixture: ComponentFixture<SongCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SongCardComponent, HttpClientModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SongCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    //set data to test
    component.post = 
    {
      user: 'Jackie',
      profile_img: 'url',
      type: 'favorite-song',
      title: 'MY FAVORITE SONG',
      content:{
        song_title: 'Engagement Party',
        song_url: 'https://open.spotify.com/track/5PYPCxyWltRIyPkhSsnWIk',
        song_embed: "https://open.spotify.com/embed/track/6LxcPUqx6noURdA5qc4BAT?utm_source=generator",
        //song_embed
      }
    };
    
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });
});
