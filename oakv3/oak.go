package main

import (
	"log"

	//"golang.org/x/exp/shiny/driver"
	//"golang.org/x/exp/shiny/screen"
	//"golang.org/x/mobile/event/lifecycle"

	// "golang.org/x/exp/shiny/imageutil"
	"github.com/oakmound/oak/v3/shiny/driver"
	"github.com/oakmound/oak/v3/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	// "golang.org/x/image/math/f64"
	// "golang.org/x/mobile/event/key"
	// "golang.org/x/mobile/event/lifecycle"
	// "golang.org/x/mobile/event/paint"
	// "golang.org/x/mobile/event/size"
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(screen.WindowGenerator{
			Title: "Basic Shiny Example",
		})
		if err != nil {
			log.Fatal(err)
		}
		defer w.Release()

		for {
			switch e := w.NextEvent().(type) {
			case lifecycle.Event:
				if e.To == lifecycle.StageDead {
					return
				}
			}
		}
	})
}
