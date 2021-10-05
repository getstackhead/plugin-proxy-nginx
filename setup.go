package main

import (
	"fmt"

	"github.com/getstackhead/pluginlib"
)

func setup() {
	fmt.Println("Setup step")
	pluginlib.StackHeadExecute("test")
}
