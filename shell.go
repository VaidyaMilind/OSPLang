package main

import (
	"OSPLang/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is OSPLang \n", user.Username)
	repl.Start(os.Stdin, os.Stdout)

}
