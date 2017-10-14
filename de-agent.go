package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	nats "github.com/nats-io/go-nats"
)

type (
	Event struct {
		Type   string
		Action string
	}
)

var natsUrl *string = flag.String("nats", "", "nats=url to your nats.. cluster?")

func main() {
	flag.Parse()
	if *natsUrl == "" {
		panic("No nats url specified.")
	}
	fmt.Printf("Sending events to %s", *natsUrl)
	rowScanner := bufio.NewScanner(os.Stdin)
	conn, err := nats.Connect(*natsUrl)

	if err != nil {
		fmt.Errorf("Error connecting to nats: ", err)
	} else {
		var event Event
		for rowScanner.Scan() {
			jsonString := rowScanner.Text()
			json.Unmarshal([]byte(jsonString), &event)

			key := fmt.Sprintf("%s.%s", event.Type, event.Action)
			err = conn.Publish(key, []byte(jsonString))

			if err != nil {
				fmt.Errorf("Publishing event threw an error: ", err)
			}
		}

		if err := rowScanner.Err(); err != nil {
			fmt.Errorf("There was an error scanning rows:", err)
		}
	}
}
