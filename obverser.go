package main

import "fmt"

type Observer interface {
	Update(Video)
}

type Subject interface {
	Attach(Observer)
	Detach(Observer)
	Notify()
}

type Video struct {
	Title string
}

type YoutubeChannel struct {
	observers []Observer
	video     Video
}

func (y *YoutubeChannel) Attach(observer Observer) {
	y.observers = append(y.observers, observer)
}

func (y *YoutubeChannel) Detach(observer Observer) {
	for i, o := range y.observers {
		if o == observer {
			y.observers = append(y.observers[:i], y.observers[i+1:]...)
			break
		}
	}
}

func (y *YoutubeChannel) Notify() {
	fmt.Println("Notifying observers...")
	fmt.Println(len(y.observers))
	for _, observer := range y.observers {
		observer.Update(y.video)
	}
}

func (y *YoutubeChannel) ReleaseNewVideo() {
	y.video = Video{Title: "Avengers"}
	y.Notify()
}

type Subscriber struct {
	name string
}

func (s Subscriber) Update(video Video) {
	fmt.Printf("%s is notified about new video: %s\n", s.name, video.Title)
}

func main() {
	var youtubeChannel Subject = &YoutubeChannel{}
	s1 := Subscriber{name: "John"}
	s2 := Subscriber{name: "Mary"}
	youtubeChannel.Attach(s1)
	youtubeChannel.Attach(s2)
	youtubeChannel.(*YoutubeChannel).ReleaseNewVideo()
}
