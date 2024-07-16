package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/danny270793/gohotreloader/command"
	"github.com/danny270793/gohotreloader/watcher"
)

func main() {
	watchPath := os.Args[1]
	executable := strings.Join(os.Args[2:], " ")

	fmt.Printf("watching path \"%s\"\nexecute \"%s\" on every file change\n", watchPath, executable)

	cmd := command.New(watchPath, executable)
	go cmd.Run()

	w := watcher.New(watchPath)

	c := make(chan string)
	go w.Read(c)
	for {
		file := <-c
		fmt.Printf("file %s changed\n", file)

		cmd.Cancel()
		go cmd.Run()
	}
}
