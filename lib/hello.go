package diffence

import "fmt"

// Hello says whatever you asks it to say
func Hello(msg string) {
	fmt.Println(noop(msg))
}

func noop(msg string) string {
	return msg
}
