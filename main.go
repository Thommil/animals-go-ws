package main

import (
	"fmt"
	"github.com/thommil/animals-go-common/model"
)

func main() {
	fmt.Printf("%s %s\n", "Helslo", model.user.Get())
}
