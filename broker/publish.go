package broker

import "fmt"

func Publish(topic, msg string) {
	fmt.Printf("published to topic - %s/ message -%s\n", topic, msg)
}
