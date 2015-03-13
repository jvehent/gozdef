gozdef - Go client to send events to MozDef
===========================================

http://godoc.org/github.com/jvehent/gozdef

```go
package main

import (
    "github.com/jvehent/gozdef"
    "log"
)

func main() {
    conf := gozdef.MqConf{
        Host:       "mozdef.rabbitmq.example.net",
        Port:       5671,
        User:       "gozdefclient",
        Pass:       "s3cr3tpassw0rd",
        Vhost:      "mozdef",
        Exchange:   "eventtask",
        RoutingKey: "eventtask",
        UseTLS:     true,
        CACertPath: "/etc/pki/CA/certs/ca.crt",
        Timeout:    "10s",
    }
    publisher, err := gozdef.InitAmqp(conf)
    if err != nil {
        log.Fatal(err)
    }

    ev, err := gozdef.NewEvent()
    if err != nil {
        log.Fatal(err)
    }
    ev.Category = "demo"
    ev.Source = "test client"
    ev.Summary = "tl;dr: everything's fine!"
    ev.Tags = append(ev.Tags, "gozdef")

    ev.Details = struct {
        SrcIP   string `json:"sourceipaddress"`
        DestIP  string `json:"destinationipaddress"`
        Offense string `json:"offense"`
        Blocked bool
    }{
        SrcIP:   "10.0.1.2",
        DestIP:  "192.168.1.5",
        Offense: "brute force",
        Blocked: true,
    }

    err = publisher.Send(ev)
    if err != nil {
        log.Fatal(err)
    }
}
```
