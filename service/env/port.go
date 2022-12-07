package env

import (
	"log"
	"os"
	"strconv"
)

var StartPort = 11400
var EndPort = 11500

func init() {
	startP := false
	start, ok := os.LookupEnv("FORK_PORT_START")
	if ok {
		atoi, err := strconv.Atoi(start)
		if err != nil {
			log.Panicln("FORK_PORT_START is not a number")
		}
		if atoi >= 65535 {
			log.Panicln("FORK_PORT_START is too large")
		}
		StartPort = atoi
		startP = true
	}
	end, ok := os.LookupEnv("FORK_PORT_END")
	if ok {
		atoi, err := strconv.Atoi(end)
		if err != nil {
			log.Panicln("FORK_PORT_END is not a number")
		}
		if atoi >= 65535 {
			log.Panicln("FORK_PORT_END is too large")
		}
		EndPort = atoi
	} else {
		if startP {
			EndPort = StartPort + 100
			if EndPort >= 65535 {
				log.Panicln("FORK_PORT_START is too large")
			}
		}
	}
}
