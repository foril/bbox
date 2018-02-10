package main

import (
	"fmt"
	"os"
	"time"

	"github.com/foril/bbox"
	"github.com/twinj/uuid"
)

var bb *bbox.Bearerbox

func main() {

	// Connect to a Bearerbox
	// Bearerbox server IP:PORT
	// Bearerbox SMSBOX-ID
	bb = &bbox.Bearerbox{Addr: "XXX.XXX.XXX.XX:13001", BoxID: "SMSBOX-ID"}
	err := bb.Connect()

	if err != nil {
		panic(fmt.Sprint(err))

	}

	fmt.Printf("Connected to %s - Box : %s", bb.Addr, bb.BoxID)

	// sending some messages
	go func() {
		time.Sleep(time.Second * 1)
		send(bb)
	}()

	// Read incoming messages
	for {
		msg, err := bb.Read()

		if err != nil {
			panic(fmt.Sprint(err))
		}

		// Check received message type from Bearerbox
		switch m := msg.(type) {

		// Is admin message ?
		case *bbox.Admin:
			admin(m)

		// Is acknowledgment message ?
		case *bbox.Ack:
			ack(m)

		// Is a short message ?
		case *bbox.Sms:
			sms(m)

		}
	}
}

func sms(s *bbox.Sms) {

	// Check short message type
	switch s.Sms_type {

	// Mo received
	case bbox.Mo:

		fmt.Println("#MO received#")
		fmt.Printf("%+v\n", s)

		// Write Ack response
		bb.Write(&bbox.Ack{bbox.Success, s.Time, s.Id})

	// DLR received => dlr-mask on sent message
	case bbox.Report_mo:

		fmt.Println("#DLR received#")
		fmt.Printf("%+v\n", s)

		// Write Ack response
		bb.Write(&bbox.Ack{bbox.Success, s.Time, s.Id})
	}
}

func admin(a *bbox.Admin) {

	fmt.Println("#Admin command received#")

	// Bearerbox shutdown
	if a.Command == bbox.Shutdown {
		fmt.Println("Bearerbox shutdown")
		os.Exit(0)
	}
}

func ack(a *bbox.Ack) {

	fmt.Println("#ACK received#")
	fmt.Printf("%+v\n", a)
}

func send(bb *bbox.Bearerbox) {

	for i := 0; i < 10; i++ {

		uid := uuid.NewV4().String() // i need an ID for this example

		sms := &bbox.Sms{
			Sms_type: bbox.Mt_push,
			Sender:   bbox.OCTSTR("GOPHER"),
			Receiver: bbox.OCTSTR("123456789"),
			Msgdata:  bbox.OCTSTR("Mesage sent from Go to Kannel bearerbox"),
			Time:     bbox.INTEGER(time.Now().Unix()),
			Smsc_id:  bbox.OCTSTR("SMPPSim"), // My smsc-id
			Id:       bbox.UUID(uid),
			Coding:   bbox.Coding_7BIT,
			Dlr_mask: bbox.INTEGER(31),
			Validity: bbox.INTEGER(time.Now().Unix() + (60 * 5)), // Validity 5 minutes
			Dlr_url:  bbox.OCTSTR("My DB ID : " + uid),           // Tip - put your own message ID, you'll can then use this ID to update your DB on receipt of the DLR.
		}

		// Send message
		bb.Write(sms)
	}
}
