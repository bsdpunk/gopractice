package main

import (
	"github.com/czwinzscher/libcaca-go"
	"log"
)

func main() {
	cv, err := caca.CreateCanvas(50, 10)
	if err != nil {
		log.Fatal("error while creating canvas: " + err.Error())
	}

	defer cv.Free()

	cv.SetColorAnsi(caca.COLOR_GREEN, caca.COLOR_BLACK) // hacker style
	cv.PutStr(0, 0, "hello")

	dp, err := caca.CreateDisplay(&cv)
	if err != nil {
		log.Fatal("error while creating display: " + err.Error())
	}

	dp.SetTitle("title")
	dp.Refresh()
	dp.GetEvent(caca.EVENT_KEY_PRESS, nil, -1)

	dp.Free()
}
