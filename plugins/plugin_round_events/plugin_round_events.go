package main

import (
	"fmt"
	structs "github.com/pavelkim/tcp_socket_server/structs"
	"log"
	"strings"
)

/*
	Example:
	ROUNDSTART	cs_assault	1102372777
	ROUNDEND	cs_assault	1119673718
*/

var PluginConfiguration structs.CommandPluginConfigurationStruct

func CommandHandlerFunction(payload string) (string, error) {

	payload_parts := strings.Split(strings.TrimRight(payload, "\n"), "\t")

	if len(payload_parts) != 3 {
		return "", fmt.Errorf("Broken payload, expected 3 pieces.")
	}

	event := payload_parts[0]
	current_map := payload_parts[1]
	map_time := payload_parts[2]

	log.Printf("%s: current_map:%s map_time:%s\n", event, current_map, map_time)
	fmt.Printf("%s: current_map:%s map_time:%s\n", event, current_map, map_time)

	if event != "ROUNDSTART" {
		message := fmt.Sprintf("%s: current_map:%s map_time:%s\n", event, current_map, map_time)
		go PluginConfiguration.Messenger.(func(string) (bool, error))(message)
	}

	response := "OK"

	return response, nil
}
