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

// Type
const (
	heartbeat INTEGER = iota
	admin
	sms
	ack
	wdp_datagram
)

// Admin command
const (
	Shutdown INTEGER = iota
	Suspend
	Resume
	Identify
	Restart
)

// Message
const (
	Mo = iota
	Mt_reply
	Mt_push
	Report_mo
	Report_mt
)

// Ack
const (
	Success = iota
	Failed
	Failed_tmp
	Buffered
)

// Coding
const (
	Coding_7BIT = iota
	Coding_8BIT
	Coding_UCS2
)
