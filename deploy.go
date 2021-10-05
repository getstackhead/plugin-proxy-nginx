package main

import (
	"fmt"

	"github.com/getstackhead/pluginlib"
)

func deploy() {
	fmt.Println("Deploy step")
	fmt.Println(pluginlib.GetProject())
	pluginlib.StackHeadExecute("test", "")
}
