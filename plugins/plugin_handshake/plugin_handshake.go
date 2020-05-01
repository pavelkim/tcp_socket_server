package main

import (
	"fmt"
	structs "github.com/pavelkim/tcp_socket_server/structs"
	"log"
)

/*
	Example:
	Hello
*/

var PluginConfiguration structs.CommandPluginConfigurationStruct

func CommandHandlerFunction(payload string) (string, error) {

	log.Printf("HELLO: answering with Hi.\n")
	fmt.Printf("HELLO: answering with Hi.\n")

	go PluginConfiguration.Messenger.(func(string) (bool, error))("HI")

	response := "Hi"

	return response, nil
}
