package main

import (
	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"log"
)

const (
	ResponseUnknownItem = "UnknwownItem"
	OpenHABFailed       = "OpenHABFailed"

	ItemSwitched = "ItemSwitched"
)

func switchHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	item, err := lookupItem(echoReq)
	if err == nil {
		actionName, _ := echoReq.GetSlotValue("Action")
		action := lookupAction(actionName)
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
	} else {
		response := lookupErrorResponse(ResponseUnknownItem)
		echoResp.OutputSpeech(response)
		log.Printf("Unknown item")
	}
}
