package main

import (
	"fmt"
	ui "github.com/gizak/termui"
	//"strings"
)

type CueList struct {
	Widget        *ui.List
	pos           int
	originalItems []string
	listItems     []string
	numItems      int
}

func NewCueList(s []string) CueList {
	o := CueList{}

	// Create list UI
	ls := ui.NewList()
	ls.ItemFgColor = ui.ColorYellow
	ls.BorderLabel = "CUES"
	ls.Height = len(s) + 2
	ls.Width = 80
	ls.Y = 0
	ls.Overflow = "hidden"

	o.Widget = ls
	o.pos = 0
	o.originalItems = s
	o.numItems = len(s)

	o.refreshCueView()
	o.Widget.Items = o.listItems

	return o
}

func (o CueList) GetSelectedCue() string {
	return o.originalItems[o.pos]
}

func (o *CueList) refreshCueView() {
	o.listItems = []string{}
	for i := 0; i < o.numItems; i++ {
		if o.pos == i {
			o.listItems = append(o.listItems, fmt.Sprintf("--> [%s](fg-bold)", o.originalItems[i]))
		} else {
			o.listItems = append(o.listItems, fmt.Sprintf("    %s", o.originalItems[i]))
		}
	}
	o.Widget.Items = o.listItems
}

func (o *CueList) Prev() {
	// Don't go out of bounds
	if o.pos == 0 {
		return
	}

	o.pos--
	par.Text = CUE_PADDING + o.GetSelectedCue()
	status.Text = "[STOPPED](fg-red)\n 00:00"
	player.Stop()
	player.SetFile(par.Text)

	o.refreshCueView()
	draw()
}

func (o *CueList) Next() {
	// Don't go out of bounds
	if o.pos == o.numItems-1 {
		return
	}

	o.pos++
	par.Text = CUE_PADDING + o.GetSelectedCue()
	status.Text = "[STOPPED](fg-red)\n 00:00"
	player.Stop()
	player.SetFile(par.Text)

	o.refreshCueView()
	draw()
}
