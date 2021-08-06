package main

import (
	"fmt"
	"os/user"
)

func main() {
	name := "anonymous"
	systemUser, err := user.Current()
	if err == nil {
		name = systemUser.Username
	}

	fmt.Printf("Hello %s! It's Ozon Go School project!\n", name)
}
