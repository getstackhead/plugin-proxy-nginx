package main

import (
	"fmt"

	"github.com/getstackhead/pluginlib"
)

func destroy() {
	fmt.Println("Destroy step")
	pluginlib.StackHeadExecute("test")
}
