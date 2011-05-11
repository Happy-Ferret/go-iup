/* 
	Copyright (C) 2011 by Jeremy Cowgar <jeremy@cowgar.com>
	
	This file is part of iup.go.

	iup.go is free software: you can redistribute it and/or modify
	it under the terms of the GNU Lesser General Public License as
	published by the Free Software Foundation, either version 3 of
	the License, or (at your option) any later version.
	
	iup.go is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.
	
	You should have received a copy of the GNU Lesser General Public
	License along with iup.go.  If not, see <http://www.gnu.org/licenses/>.
*/

package iup

/*
#include <stdlib.h>
#include <iup.h>
*/
import "C"
import "unsafe"

func (ih *Ihandle) StoreAttribute(name, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	
	C.IupStoreAttribute(ih.h, cName, cValue)
}

func (ih *Ihandle) StoreAttributeId(name string, id int, value string) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	cValue := C.CString(value)
	defer C.free(unsafe.Pointer(cValue))
	
	C.IupStoreAttributeId(ih.h, cName, C.int(id), cValue)
}

/*

func (ih *Ihandle) SetAttribute(name string, value unsafe.Pointer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	C.IupSetAttribute(ih.h, cName, value)
}

func (ih *Ihandle) SetAttributeId(name string, id int, value unsafe.Pointer) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	C.IupSetAttributeId(ih.h, cName, C.int(id), value)
}
*/
/*
func (ih *Ihandle) SetfAttribute(name, format string, args ...interface{}) {
	ih.StoreAttribute(name, fmt.Sprintf(format, args...))
}

func (ih *Ihandle) SetfAttributeId(name string, id int, format string, args ...interface{}) {
	ih.StoreAttributeId(name, id, fmt.Sprintf(format, args...))
}

func (ih *Ihandle) SetfAttributeId2(name string, lin, col int, format string, args ...interface{}) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	cValue := C.CString(fmt.Sprintf(format, args...))
	defer C.free(unsafe.Pointer(cValue))
	
	C.IupSetfAttributeId2(ih.h, cName, C.int(lin), C.int(col), cValue)
}
*/
func (ih *Ihandle) SetAttributes(values string) {
	cValues := C.CString(values)
	defer C.free(unsafe.Pointer(cValues))
	
	C.IupSetAttributes(ih.h, cValues)
}

func (ih *Ihandle) SetAttributeHandle(name string, ihNamed *Ihandle) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	C.IupSetAttributeHandle(ih.h, cName, ihNamed.h)
}

func (ih *Ihandle) GetAttribute(name string) string {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	cVal := C.IupGetAttribute(ih.h, cName)
	
	return C.GoString(cVal)
}

func (ih *Ihandle) GetFloat(name string) float64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	return float64(C.IupGetFloat(ih.h, cName))
}

func (ih *Ihandle) GetInt(name string) int64 {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	
	return int64(C.IupGetInt(ih.h, cName))
}
