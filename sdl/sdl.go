package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/damoon/lab-raytracing/raytracing"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	err := Window()
	if err != nil {
		log.Fatal(err)
	}
}

func Window() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		return fmt.Errorf("could not initialize SDL: %v", err)
	}
	defer sdl.Quit()

	w, err := sdl.CreateWindow("winTitle", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		1280, 720, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("create window: %v", err)
	}
	defer w.Destroy()

	r, err := sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return fmt.Errorf("create renderer: %v", err)
	}
	defer r.Destroy()

	s, err := raytracing.NewScene(r)
	if err != nil {
		return fmt.Errorf("create scene: %v", err)
	}

	log.Print("create circle")
	go raytracing.Circle(s.Img)
	log.Print("create circle done")

	events := make(chan sdl.Event)
	errc := s.Run(events, r)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case err := <-errc:
			return err
		}
	}
}
