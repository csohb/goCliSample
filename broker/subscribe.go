package broker

import (
	"fmt"
	"time"
)

func Subscribe(topic string) {
	for {
		fmt.Printf("subscribe topic : %s\n", topic)
		time.Sleep(time.Second * 1)
	}
}
