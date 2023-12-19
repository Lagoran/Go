package message

import (
	"fmt"
)

func Greeting(salute, name string) (salutation string) {
	salutation = fmt.Sprintf("%s, %s", salute, name)
	return
}
