package gst

/*
#include <stdlib.h>
#include <gst/gst.h>
*/
import "C"

import (
	"unsafe"
	"github.com/ziutek/glib"
)

type PadLinkReturn C.GstPadLinkReturn

const (
  PAD_LINK_OK = PadLinkReturn(C.GST_PAD_LINK_OK)
  PAD_LINK_WRONG_HIERARCHY = PadLinkReturn(C.GST_PAD_LINK_WRONG_HIERARCHY)
  PAD_LINK_WAS_LINKED = PadLinkReturn(C.GST_PAD_LINK_WAS_LINKED)
  PAD_LINK_WRONG_DIRECTION = PadLinkReturn(C.GST_PAD_LINK_WRONG_DIRECTION)
  PAD_LINK_NOFORMAT = PadLinkReturn(C.GST_PAD_LINK_NOFORMAT)
  PAD_LINK_NOSCHED = PadLinkReturn(C.GST_PAD_LINK_NOSCHED)
  PAD_LINK_REFUSED = PadLinkReturn(C.GST_PAD_LINK_REFUSED)
)

func (p PadLinkReturn) String() string {
	switch p {
	case PAD_LINK_OK:
		return "PAD_LINK_OK"
	case PAD_LINK_WRONG_HIERARCHY:
		return "PAD_LINK_WRONG_HIERARCHY"
	case PAD_LINK_WAS_LINKED:
		return "PAD_LINK_WAS_LINKED"
	case PAD_LINK_WRONG_DIRECTION:
		return "PAD_LINK_WRONG_DIRECTION"
	case PAD_LINK_NOFORMAT:
		return "PAD_LINK_NOFORMAT"
	case PAD_LINK_NOSCHED:
		return "PAD_LINK_NOSCHED"
	case PAD_LINK_REFUSED:
		return "PAD_LINK_REFUSED"
	}
	panic("Wrong value of PadLinkReturn variable")
}

type PadDirection C.GstPadDirection

const (
  PAD_UNKNOWN = PadDirection(C.GST_PAD_UNKNOWN)
  PAD_SRC = PadDirection(C.GST_PAD_SRC)
  PAD_SINK = PadDirection(C.GST_PAD_SINK)
)

func (p PadDirection) g() C.GstPadDirection {
	return C.GstPadDirection(p)
}

func (p PadDirection) String() string {
	switch p {
	case PAD_UNKNOWN:
		return "PAD_UNKNOWN"
	case PAD_SRC:
		return "PAD_SRC"
	case PAD_SINK:
		return "PAD_SINK"
	}
	panic("Wrong value of PadDirection variable")
}

type Pad struct {
	GstObj
}

func (p *Pad) g() *C.GstPad {
	return (*C.GstPad)(p.GetPtr())
}

func (p *Pad) AsPad() *Pad {
	return p
}

func (p *Pad) CanLink(sink_pad *Pad) bool {
	return C.gst_pad_can_link(p.g(), sink_pad.g()) != 0
}

func (p *Pad) Link(sink_pad *Pad) PadLinkReturn {
	return PadLinkReturn(C.gst_pad_link(p.g(), sink_pad.g()))
}

func (p *Pad) GetCurrentCaps() *Caps {
	// unref the caps when you're done
	return (*Caps)( C.gst_pad_get_current_caps((*C.GstPad)(p.GetPtr())) )
}

type GhostPad struct {
	Pad
}

func (p *GhostPad) g() *C.GstGhostPad {
	return (*C.GstGhostPad)(p.GetPtr())
}

func (p *GhostPad) AsGhostPad() *GhostPad {
	return p
}

func (p *GhostPad) SetTarget(new_target *Pad) bool {
	return C.gst_ghost_pad_set_target(p.g(), new_target.g()) != 0
}

func (p *GhostPad) GetTarget() *Pad {
	r := new(Pad)
	r.SetPtr(glib.Pointer(C.gst_ghost_pad_get_target(p.g())))
	return r
}

func (p *GhostPad) Construct() bool {
	return C.gst_ghost_pad_construct(p.g()) != 0
}

func NewGhostPad(name string, target *Pad) *GhostPad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(GhostPad)
	p.SetPtr(glib.Pointer(C.gst_ghost_pad_new(s, target.g())))
	return p
}

func NewGhostPadNoTarget(name string, dir PadDirection) *GhostPad {
	s := (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(s))
	p := new(GhostPad)
	p.SetPtr(glib.Pointer(C.gst_ghost_pad_new_no_target(s, dir.g())))
	return p
}
