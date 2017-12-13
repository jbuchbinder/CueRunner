package main

import (
	"fmt"
	"github.com/ziutek/gst"
	"path/filepath"
	"time"
)

type Player struct {
	player       *gst.Element
	playing      bool
	playduration int
}

func NewPlayer() *Player {
	obj := &Player{}
	p := gst.ElementFactoryMake("playbin", "player")
	obj.player = p
	obj.playing = false

	return obj
}

func (p *Player) SetFile(fn string) {
	abspath, _ := filepath.Abs(fn)
	p.player.SetProperty("uri", "file://"+abspath)
	p.playduration = 0
}

func (p *Player) Play() {
	status.Text = "PLAYING\n" + p.GetTime()
	draw()
	p.player.SetState(gst.STATE_PLAYING)
	p.playing = true
	go func(p *Player) {
		for p.playing == true {
			p.playduration += 1
			time.Sleep(time.Duration(1) * time.Second)
			status.Text = "PLAYING\n" + p.GetTime()
			draw()
		}
		status.Text = "STOPPED"
	}(p)
}

func (p *Player) Pause() {
	status.Text = "PAUSED\n" + p.GetTime()
	draw()
	p.player.SetState(gst.STATE_PAUSED)
	p.playing = false
}

func (p *Player) Stop() {
	status.Text = "STOPPED\n" + p.GetTime()
	draw()
	p.player.SetState(gst.STATE_NULL)
	p.playing = false
	p.playduration = 0
}

func (p *Player) IsPlaying() bool {
	return p.playing
}

func (p *Player) GetTime() string {
	x := p.playduration
	if x == 0 {
		return "00:00"
	}
	return fmt.Sprintf("%02d:%02d", x/60, x%60)
}
