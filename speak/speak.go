package speak

import (
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"log"
	"os"
	"time"
)
func SayFail(){
	stream,_ := os.Open("1.wav")
	streamer,format,err := wav.Decode(stream)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	time.Sleep(time.Second * 10)
}
func SayOpen(){
	stream,_ := os.Open("2.wav")
	streamer,format,err := wav.Decode(stream)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	time.Sleep(time.Second * 10)
}
func SayNot(){
	stream,_ := os.Open("3.wav")
	streamer,format,err := wav.Decode(stream)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	time.Sleep(time.Second * 10)
}
func SayNoRfid(){
	stream,_ := os.Open("4.wav")
	streamer,format,err := wav.Decode(stream)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)
	time.Sleep(time.Second * 10)
}