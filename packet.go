/*

MIT License

Copyright (c) 2018 foril

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

package bbox

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
)

// Decode io.reader
func decode(r io.Reader) (interface{}, error) {

	h := make([]byte, 0x4)

	n, err := io.ReadFull(r, h)
	l := binary.BigEndian.Uint32(h) // Size

	// TODO improve error message
	if uint32(n) < 4 || err != nil {
		return nil, err
	}

	b := make([]byte, l)
	n, err = io.ReadFull(r, b)

	// TODO improve error message
	if uint32(n) < l || err != nil {
		return nil, err
	}

	c := binary.BigEndian.Uint32(b[:4]) // Cmd
	b = b[4:]                           // Data

	var i interface{}

	switch INTEGER(c) {
	case heartbeat:
		i = new(Heartbeat)
	case admin:
		i = new(Admin)
	case sms:
		i = new(Sms)
	case ack:
		i = new(Ack)
	default:
		return nil, fmt.Errorf("Unkown message received \n ################### \n Length : %v \n Command : %v \n Data : \n %v \n ################### \n", l, c, b)
	}

	return unpack(i, b, l), nil
}

// Unpack message to struct
func unpack(i interface{}, b []byte, l uint32) interface{} {

	var f = reflect.ValueOf(i).Elem()

	for i := 0; i < f.NumField(); i++ {

		if f.Field(i).Type().Name() == "OCTSTR" || f.Field(i).Type().Name() == "UUID" {
			cl := binary.BigEndian.Uint32(b[:4])
			if cl <= l {
				b = b[4:]
				f.Field(i).SetString(string(b[0:cl]))
				b = b[cl:]
			} else {
				b = b[4:]
			}

		} else if f.Field(i).Type().Name() == "INTEGER" {
			f.Field(i).SetInt(int64(binary.BigEndian.Uint32(b[:4])))
			b = b[4:]
		}
	}

	return i
}

// Encode message
func encode(i interface{}) *bytes.Buffer {

	var b bytes.Buffer

	switch i.(type) {
	case *Ack:
		b.Write(ack.INTEGER())
	case *Admin:
		b.Write(admin.INTEGER())
	case *Heartbeat:
		b.Write(heartbeat.INTEGER())
	case *Sms:
		b.Write(sms.INTEGER())
	}

	return pack(b, i)
}

// Pack struct to message
func pack(b bytes.Buffer, i interface{}) *bytes.Buffer {

	var f = reflect.ValueOf(i).Elem()

	for i := 0; i < f.NumField(); i++ {

		switch f.Field(i).Type().Name() {
		case "OCTSTR":
			b.Write(OCTSTR(f.Field(i).String()).OCTSTR())
		case "UUID":
			b.Write(UUID(f.Field(i).String()).UUID())
		case "INTEGER":
			b.Write(INTEGER(f.Field(i).Int()).INTEGER())
		}
	}

	m := new(bytes.Buffer)
	binary.Write(m, binary.BigEndian, int32(len(b.Bytes()))) // Full length
	binary.Write(m, binary.BigEndian, b.Bytes())             // Message

	return m

}
