package cmd

import (
	"fmt"

	"github.com/captainviet/testgopls/constants"
)

// Hello prints the Unit value of 20
func Hello() {
	fmt.Println(constants.Unit(20))
}
