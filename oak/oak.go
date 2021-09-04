package main

import (
	"log"

	//"golang.org/x/exp/shiny/driver"
	//"golang.org/x/exp/shiny/screen"
	//"golang.org/x/mobile/event/lifecycle"

	// "golang.org/x/exp/shiny/imageutil"
	"github.com/oakmound/shiny/driver"
	"github.com/oakmound/shiny/screen"
	"golang.org/x/mobile/event/lifecycle"
	// "golang.org/x/image/math/f64"
	// "golang.org/x/mobile/event/key"
	// "golang.org/x/mobile/event/lifecycle"
	// "golang.org/x/mobile/event/paint"
	// "golang.org/x/mobile/event/size"
)

func main() {
	driver.Main(func(s screen.Screen) {
		w, err := s.NewWindow(screen.WindowGenerator{})
		if err != nil {
			log.Print(err)
			return
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
