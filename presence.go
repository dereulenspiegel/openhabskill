package main

import (
	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"log"
)

const (
	Goodbye = "Goodbye"
	Hello   = "Hello"
)

func presenceHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse, away bool) {
	item, err := lookupItem(echoReq)
	if err != nil {
		response := lookupErrorResponse(ResponseUnknownItem)
		echoResp.OutputSpeech(response)
		log.Printf("Unknown item")
		return
	}
	var action string
	if away {
		action = "OFF"
	} else {
		action = "ON"
	}
	log.Printf("Sending %s for item %s", action, item)
	err = oh.SendCommand(item, action)
	if err != nil {
		response := lookupErrorResponse(OpenHABFailed)
		echoResp.OutputSpeech(response)
		log.Printf("Failed to send command to OpenHAB: %+v\n", err)
	} else {
		response := lookupResponse(ItemSwitched)
		echoResp.OutputSpeech(response)
	}
}
