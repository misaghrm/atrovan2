package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	rw sync.RWMutex
)

func main() {
	var a, b string
	ch := make(chan string)
	timer := time.NewTicker(5 * time.Second)

	for {
		go func() {
			d := <-ch
			if len(d) != 0 {
				rw.Lock()
				a = ""
				rw.Unlock()
				fmt.Println(d[len(d)-1:])
			}
		}()
		go func() {
			for {
				select {
				case <-timer.C:
					ch <- a
				default:
				}
			}
		}()

		fmt.Scanln(&b)
		if strings.Contains(b, "e") {
			os.Exit(1)
		}
		b = strings.Replace(b, "\n", "", -1)
		b = strings.Replace(b, " ", "", -1)
		a = a + b
		b = ""
	}
}
