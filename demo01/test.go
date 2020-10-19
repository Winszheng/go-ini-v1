package main

import (
	"fmt"
	myini "github.com/Winszheng/go-ini-v1"
	"os"
)

func main() {
	filename := "demo01/my.ini"
	f, err := myini.Load(filename)
	if err!= nil {
		fmt.Println("failed to read file")
		os.Exit(1)
	}

	//典型的读取操作
	fmt.Println("App Mode: ", f.Section("").Key("app_mode"))
	fmt.Println("Data Path: ", f.Section("paths").Key("data"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("protocol"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("http_port"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("enforce_domain"))


	var listen myini.List
	f, err = myini.Watch(filename, listen)
	if err!= nil {
		fmt.Println("failed to read file")
		os.Exit(1)
	}
	fmt.Println()
	fmt.Println(filename, " has been changed, here are newest version: ")
	fmt.Println()
	//典型的读取操作
	fmt.Println("App Mode: ", f.Section("").Key("app_mode"))
	fmt.Println("Data Path: ", f.Section("paths").Key("data"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("protocol"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("http_port"))
	fmt.Println("Server Protocol: ", f.Section("server").Key("enforce_domain"))
}
