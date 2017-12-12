package main

import (
	//"fmt"
	ui "github.com/gizak/termui"
	"io/ioutil"
	"strings"
)

var (
	draw              func()
	par, status, info *ui.Par
	player            *Player
)

func main() {
	// Import cues from cues file
	cuelist, err := ioutil.ReadFile("cues.txt")
	if err != nil {
		panic(err)
	}
	cues := strings.Split(string(cuelist), "\n")

	// Trim last entry if empty
	if cues[len(cues)-1] == "" {
		cues = cues[:len(cues)-2]
	}

	//fmt.Printf("%#v\n", cues)

	err = ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	player = NewPlayer()

	par = ui.NewPar(cues[0])
	par.Height = 3
	par.Width = 50
	par.TextFgColor = ui.ColorWhite
	par.BorderLabel = "CURRENT CUE"
	par.BorderFg = ui.ColorCyan

	status = ui.NewPar("STOPPED")
	status.Height = 3
	status.Width = 30
	status.TextFgColor = ui.ColorWhite
	status.BorderLabel = "STATE"
	status.BorderFg = ui.ColorCyan

	info = ui.NewPar("Ctrl-X : Exit   |   Enter : Play/Pause   |  Up/Down : Navigate")
	info.Height = 4
	info.Width = 80
	info.TextFgColor = ui.ColorWhite
	info.BorderLabel = "INFO"
	info.BorderFg = ui.ColorWhite

	// Create list UI
	cl := NewCueList(cues)

	ui.Body.AddRows(
		ui.NewRow(
			ui.NewCol(12, 0, par),
		),
		ui.NewRow(
			ui.NewCol(9, 0, cl.Widget),
			ui.NewCol(3, 0, status),
		),
		ui.NewRow(
			ui.NewCol(12, 0, info),
		),
	)

	draw = func() {
		ui.Render(ui.Body)
	}

	// handle key q pressing
	ui.Handle("/sys/kbd/C-x", func(ui.Event) {
		// handle Ctrl + x combination
		ui.StopLoop()
	})
	//ui.Handle("/sys/kbd", func(ui.Event) {
	//	// handle all other key pressing
	//})
	// List selection
	ui.Handle("/sys/kbd/<up>", func(ui.Event) {
		cl.Prev()
	})
	ui.Handle("/sys/kbd/<down>", func(ui.Event) {
		cl.Next()
	})
	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		if player.IsPlaying() {
			player.Pause()
		} else {
			player.Play()
		}
	})

	draw()
	ui.Body.Align()

	ui.Render(ui.Body)

	ui.Loop()
}
