/*

MIT License

Copyright (c) 2018 Farid TOUIL - touilf@gmail.com

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
)

type INTEGER int32

type OCTSTR string

type VOID string

type UUID string

type Heartbeat struct {
	Load INTEGER
}

type Admin struct {
	Command INTEGER
	Boxc_id OCTSTR
}

type Sms struct {
	Sender      OCTSTR
	Receiver    OCTSTR
	Udhdata     OCTSTR
	Msgdata     OCTSTR
	Time        INTEGER
	Smsc_id     OCTSTR
	Smsc_number OCTSTR
	Foreign_id  OCTSTR
	Service     OCTSTR
	Account     OCTSTR // Max length 64
	Id          UUID
	Sms_type    INTEGER
	Mclass      INTEGER
	Mwi         INTEGER
	Coding      INTEGER
	Compress    INTEGER
	Validity    INTEGER
	Deferred    INTEGER
	Dlr_mask    INTEGER
	Dlr_url     OCTSTR
	Pid         INTEGER
	Alt_dcs     INTEGER
	Rpi         INTEGER
	Charset     OCTSTR
	Boxc_id     OCTSTR
	Binfo       OCTSTR
	Msg_left    INTEGER
	Split_parts VOID // VOID
	Priority    INTEGER
	Resend_try  INTEGER
	Resend_time INTEGER
	Meta_data   OCTSTR
}

type Ack struct {
	Nack INTEGER
	Time INTEGER
	Id   UUID
}

type Wdp_datagram struct {
	Source_address      OCTSTR
	Source_port         INTEGER
	Destination_address OCTSTR
	Destination_port    INTEGER
	User_data           OCTSTR
}

func (i INTEGER) INTEGER() []byte {

	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, i)

	return b.Bytes()
}

func (o OCTSTR) OCTSTR() []byte {

	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, int32(len(o)))

	return append(b.Bytes(), []byte(o)...)
}

func (u UUID) UUID() []byte {

	b := new(bytes.Buffer)
	binary.Write(b, binary.BigEndian, int32(len(u)))

	return append(b.Bytes(), []byte(u)...)
}
