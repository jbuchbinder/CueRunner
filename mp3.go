package main

import (
	"github.com/ziutek/gst"
	"path/filepath"
)

type Player struct {
	player  *gst.Element
	playing bool
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
}

func (p *Player) Play() {
	status.Text = "PLAYING"
	draw()
	p.player.SetState(gst.STATE_PLAYING)
	p.playing = true
}

func (p *Player) Pause() {
	status.Text = "PAUSE"
	draw()
	p.player.SetState(gst.STATE_PAUSED)
	p.playing = false
}

func (p *Player) Stop() {
	status.Text = "STOPPED"
	draw()
	p.player.SetState(gst.STATE_NULL)
	p.playing = false
}

func (p *Player) IsPlaying() bool {
	return p.playing
}
