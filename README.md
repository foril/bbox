# bbox
[![Build Status](https://travis-ci.org/foril/bbox.svg?branch=master)](https://travis-ci.org/foril/bbox) [![Go Report Card](https://goreportcard.com/badge/github.com/foril/bbox)](https://goreportcard.com/report/github.com/foril/bbox)

Bbox is a Go library for Kannel SMS Gateway Box protocol
*(Based on Kannel 1.4.4)*

Create your own "Custom BOX" for send, receive SMS and handle DLR easily with Kannel Bearerbox .

### Install
```
go get github.com/foril/bbox
```
## Sample Box

The following example can be found in [samples](https://github.com/foril/bbox/tree/master/samples).

### Connect to the Bearerbox

```Go
bb = &bbox.Bearerbox{Addr: "XXX.XXX.XXX.XX:13001", BoxID: "SMSBOX-ID"}
err := bb.Connect()

if err != nil {
  panic(fmt.Sprint(err))
}

// Maintain persistent connection
for {
  msg, err := bb.Read()
  
  ....
 }
...
```

### Receiving messages from the Bearerbox

```Go
// Reading all incoming messages
for {
  msg, err := bb.Read()

  if err != nil {
    panic(fmt.Sprint(err))
  }

  // Check received message type from Bearerbox
  switch m := msg.(type) {

  // Is admin message ?
  case *bbox.Admin:
    // Do stuff

  // Is acknowledgment message ?
  case *bbox.Ack:
    // Do stuff

  // Is a short message ?
  case *bbox.Sms:
    // Do stuff
  }
}
```

### Sending SMS to the Bearerbox

```Go
sms := &bbox.Sms{
  Sms_type: bbox.Mt_push,
  Sender:   bbox.OCTSTR("GOPHER"),
  Receiver: bbox.OCTSTR("123456789"),
  Msgdata:  bbox.OCTSTR("Mesage sent from Go to Kannel bearerbox"),
  Time:     bbox.INTEGER(time.Now().Unix()),
  Smsc_id:  bbox.OCTSTR("SMPPSim"), // My smsc-id
  Id:       bbox.UUID(uid), // You can use an extrenal library for generate UUID
  Coding:   bbox.Coding_7BIT,
  Dlr_mask: bbox.INTEGER(31),
  Validity: bbox.INTEGER(time.Now().Unix() + (60 * 5)), // Validity 5 minutes
  Dlr_url:  bbox.OCTSTR("My DB ID : " + uid), // Tip - put your own message ID, you'll can then use this ID to update your DB on receipt of the DLR.
}

// Send message
bb.Write(sms)
```

### Filter DLR / MO on messages received from the Bearerbox

```Go
// Check short message type *bbox.Sms.Sms_type
switch sms.Sms_type {

// Mo received
case bbox.Mo:

  fmt.Println("#MO received#")
  // Do stuff

  // Write Ack response
  bb.Write(&bbox.Ack{bbox.Success, sms.Time, sms.Id})

// DLR received => dlr-mask on sent message
case bbox.Report_mo:

  fmt.Println("#DLR received#")
  // Do stuff

  // Write Ack response
  bb.Write(&bbox.Ack{bbox.Success, sms.Time, sms.Id})
}
```

## You have built an custom box [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/foril/bbox/issues)
Feel free to contribute.
