package reporter

import(
	"../go-hpfeeds/"
	"fmt"
	"time"
)

func StartReporter() {

	host := "http://159.203.253.60"
	port := 20000
	ident := "1ca0ec5a-eb51-11e5-a612-0401af4dc501"
	auth := "L1uUI9Tw"

	hp := hpfeeds.NewHpfeeds(host, port, ident, auth)
	hp.Log = true
	hp.Connect()

	// Publish something on "flotest" every second
	// Channel1 is where we put the filtered JSON data
	channel1 := make(chan []byte)
	hp.Publish("beeswax.events", channel1)
	go func() {
		for {
			channel1 <- []byte("Something")
			time.Sleep(time.Second)
		}
	}()

	// Subscribe to "flotest" and print everything coming in on it
	// prints something once every second - verify with others ::
	channel2 := make(chan hpfeeds.Message)
	hp.Subscribe("beeswax.events", channel2)
	go func() {
		for foo := range channel2 {
			fmt.Println(foo.Name, string(foo.Payload))
		}
	}()

	// Wait for disconnect
	<-hp.Disconnected
}