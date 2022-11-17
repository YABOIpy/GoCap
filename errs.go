package gocap

import (
	"log"
)

func call_err(Err error, mode bool) {
	if mode == true {
		for {
			if Err != nil {
				log.Fatal(Err)
				continue
			}
		}
	} else if mode == false {
		if Err != nil {
			log.Fatal(Err)
		}
	}
}