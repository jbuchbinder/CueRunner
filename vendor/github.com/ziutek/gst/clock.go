package gst

/*
#include <gst/gst.h>
*/
import "C"

type Clock struct {
	GstObj
}

func (c *Clock) g() *C.GstClock {
	return (*C.GstClock)(c.GetPtr())
}

func (c *Clock) AsClock() *Clock {
	return c
}
