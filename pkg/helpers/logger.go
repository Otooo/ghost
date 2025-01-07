package helpers

import (
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %v", msg, err)
	}
}

func SentConfirm(msg string) {
	log.Printf(" [x] Sent %s\n", msg)
}

func Message(msg string) {
	log.Printf(" Message is : %s\n", msg)
}

func FatalError(msg string) {
	log.Fatal(msg)
}
