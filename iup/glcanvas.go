/* 
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>

	This file is part of go-iup.

	go-iup is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.

	go-iup is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU Lesser General Public
	License along with go-iup.  If not, see <http://www.gnu.org/licenses/>.
*/

package iup

/*
#include <stdlib.h>
#include <iup.h>
#include <iupgl.h>
*/
import "C"

func GLCanvas(opts ...interface{}) *Ihandle {
	ensureControlLibOpened()

	ih := &Ihandle{h: C.IupGLCanvas(nil)}

	for _, o := range opts {
		switch v := o.(type) {
		case string:
			ih.SetAttributes(v)
		}
	}

	return ih
}

func (ih *Ihandle) GLMakeCurrent() {
	C.IupGLMakeCurrent(ih.h)
}

func (ih *Ihandle) GLIsCurrent() int {
	return int(C.IupGLIsCurrent(ih.h))
}

func (ih *Ihandle) GLSwapBuffers() {
	C.IupGLSwapBuffers(ih.h)
}

func (ih *Ihandle) GLPalette(index int, r, g, b float64) {
	C.IupGLPalette(ih.h, C.int(index), C.float(r), C.float(g), C.float(b))
}

func (ih *Ihandle) GLUseFont(first, count, list_base int) {
	C.IupGLUseFont(ih.h, C.int(first), C.int(count), C.int(list_base))
}

func GLWait(gl int) {
	C.IupGLWait(C.int(gl))
}
