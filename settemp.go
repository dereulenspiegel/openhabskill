package main

import (
	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"log"
)

func settempHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	echoReq.Request.Intent.Slots["ItemName"] = alexa.EchoSlot{
		Name:  "ItemName",
		Value: "heizung",
	}
	item, err := lookupItem(echoReq)
	if err != nil {
		response := lookupErrorResponse(ResponseUnknownItem)
		echoResp.OutputSpeech(response)
		log.Printf("Unknown item")
		return
	}
	degree, _ := echoReq.GetSlotValue("Degree")
	log.Printf("Sending %s for item %s", degree, item)
	err = oh.SendCommand(item, degree)
	if err != nil {
		response := lookupErrorResponse(OpenHABFailed)
		echoResp.OutputSpeech(response)
		log.Printf("Failed to send command to OpenHAB: %+v\n", err)
	} else {
		response := lookupResponse(ItemSwitched)
		echoResp.OutputSpeech(response)
	}
}
