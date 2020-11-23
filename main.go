package main

import (
	"flag"
	//"fmt"
	"io/ioutil"
	"strings"

	ui "github.com/gizak/termui"
)

const (
	CUE_PADDING = "          "
)

var (
	CuesFile = flag.String("cues-file", "cues.txt", "Filename containing cues")

	draw              func()
	par, status, info *ui.Par
	player            *Player
)

func main() {
	flag.Parse()

	// Import cues from cues file
	cuelist, err := ioutil.ReadFile(*CuesFile)
	if err != nil {
		panic(err)
	}
	cues := strings.Split(string(cuelist), "\n")

	// Trim last entry if empty
	if cues[len(cues)-1] == "" {
		cues = cues[:len(cues)-2]
	}

	err = ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	player = NewPlayer()

	par = ui.NewPar(CUE_PADDING + cues[0])
	par.Height = 3
	par.Width = 50
	par.TextFgColor = ui.ColorWhite
	par.BorderLabel = "CURRENT CUE"
	par.BorderFg = ui.ColorCyan

	status = ui.NewPar("[STOPPED](fg-red)\n 00:00")
	status.Height = 4
	status.Width = 30
	status.TextFgColor = ui.ColorWhite
	status.BorderLabel = "STATE"
	status.BorderFg = ui.ColorCyan

	info = ui.NewPar("  [Ctrl-X](fg-bold) : Exit   |   [Enter/Space](fg-bold) : Play/Pause   |  [Up/Down](fg-bold) : Navigate")
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
	ui.Handle("/sys/kbd/<space>", func(ui.Event) {
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
