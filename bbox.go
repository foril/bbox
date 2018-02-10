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
	"bufio"
	"io"
	"net"
	"sync"
)

// Create a Bearerbox client
// Addr 127.0.0.1:13001
// BoxId kannel smsbox-id
type Bearerbox struct {
	Addr  string
	BoxID string
	conn
}

type conn struct {
	rwc net.Conn
	r   *bufio.Reader
	w   *bufio.Writer
	m   sync.Mutex
}

// Connect the client to the Bearerbox
func (b *Bearerbox) Connect() error {

	s, err := net.Dial("tcp", b.Addr)

	// TODO improve error message
	if err != nil {
		return err
	}

	b.rwc = s
	b.r = bufio.NewReader(s)
	b.w = bufio.NewWriter(s)

	err = b.Write(&Admin{INTEGER(Identify), OCTSTR(b.BoxID)})

	// TODO improve error message
	if err != nil {
		return err
	}

	return nil
}

// Write message to the Bearerbox
func (c *conn) Write(m interface{}) error {
	c.m.Lock()
	defer c.m.Unlock()
	_, err := io.Copy(c.w, encode(m))

	// TODO improve error message
	if err != nil {
		return err
	}

	err = c.w.Flush()

	// TODO improve error message
	if err != nil {
		return err
	}

	return nil
}

// Read message from the Bearerbox
func (c *conn) Read() (interface{}, error) {

	msg, err := decode(c.r)

	// TODO improve error message
	if err != nil {
		return nil, err
	}

	return msg, nil
}

// Close the current client connection
func (c *conn) Close() error {

	err := c.rwc.Close()

	// TODO improve error message
	if err != nil {
		return err
	}

	return nil
}
