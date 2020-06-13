package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mut sync.Mutex
var smsg string = "tick"
var mmsg string = "tock"
var hmsg string = "bong"
var msg string

var gque = make(chan string, 10)

func main() {

	wg.Add(5)

	go print()
	go tick()
	go tock()
	go bong()
	go input()

	wg.Wait()

}

func print() {
	for {

		fmt.Println(time.Now(), <-gque)
		time.Sleep(time.Second)

	}
}

func tick() {
	for {
		time.Sleep(time.Second)

		mut.Lock()

		gque <- smsg
		mut.Unlock()

	}
	wg.Done() //no need for this
}

func tock() {
	for {

		time.Sleep(time.Minute)
		mut.Lock()
		gque <- mmsg
		mut.Unlock()

	}
	wg.Done() //no need for this
}

func bong() {
	hr := 0
	for {
		time.Sleep(time.Hour)
		mut.Lock()
		gque <- hmsg
		mut.Unlock()

		hr++
		if hr > 2 {
			//Exit after 3 hrs
			os.Exit(0)
		}

	}
	wg.Done() //no need for this

}

func input() {

	for {
		time.Sleep(time.Microsecond)
		sc := bufio.NewScanner(os.Stdin)

		sc.Scan()

		msg = sc.Text()

		cmd := strings.Split(msg, " ")

		if cmd == nil || len(cmd) != 2 {
			fmt.Println("Syntax error :  s/m/h newstring")
		}

		mut.Lock()

		if cmd[0] == "s" {
			smsg = cmd[1]
		}
		if cmd[0] == "m" {
			mmsg = cmd[1]

		}
		if cmd[0] == "h" {
			hmsg = cmd[1]

		}
		mut.Unlock()
	}
}
