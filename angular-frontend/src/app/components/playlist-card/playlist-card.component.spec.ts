import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientModule } from '@angular/common/http';
import { PlaylistCardComponent } from './playlist-card.component';

describe('PlaylistCardComponent', () => {
  let component: PlaylistCardComponent;
  let fixture: ComponentFixture<PlaylistCardComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PlaylistCardComponent, HttpClientModule]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PlaylistCardComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    //set data to test
    component.post = 
    {
      user: 'Gaby',
      profile_img: 'url',
      type: 'playlist',
      title: 'MY NEW PLAYLIST',
      content:{
        playlist_title: "Study playlist",
        playlist_url: "https://open.spotify.com/playlist/1yJb4XCnM4KfeO2UkMAYnp?si=945e25fe87034d38",
        playlist_embed: "https://open.spotify.com/embed/playlist/1yJb4XCnM4KfeO2UkMAYnp?utm_source=generator",
      }
    };
    
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });
});
