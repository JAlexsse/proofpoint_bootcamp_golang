package mylogger

import "log"

func ListerForLog(ch chan string) {
	for {
		msg := ch
		log.Println(msg)
	}
}
