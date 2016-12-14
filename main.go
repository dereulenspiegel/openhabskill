package main

import (
	"github.com/dereulenspiegel/openhab-cli/openhab"
	alexa "github.com/mikeflynn/go-alexa/skillserver"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"flag"
	"io/ioutil"
	"log"
)

const (
	UnknownCommand = "UnknownCommand"
)

var (
	cfgFilePath = flag.String("config", "", "Path to configuration file")
)

var (
	ItemNotFound = errors.New("Item not found")
)

type Config struct {
	ApplicationId   string
	UserId          string
	OpenHAB         string
	Greeting        string
	Items           map[string]map[string]string
	Modes           map[string]map[string]string
	Metrics         map[string]map[string]string
	Actions         map[string]string
	ErrorResponses  map[string]string
	SuccessResponse map[string]string
}

var defaultErrorResponse = "Ich kann das leider nicht tun"
var defaultSuccessResponse = "Sofort"

var config Config

var oh *openhab.Client

func main() {
	flag.Parse()

	cfgBytes, err := ioutil.ReadFile(*cfgFilePath)
	if err != nil {
		log.Fatalf("Can't read configuration file: %+v", err)
	}
	err = yaml.Unmarshal(cfgBytes, &config)
	if err != nil {
		log.Fatalf("Can't parse configuration file: %+v", err)
	}

	oh = openhab.NewClient(config.OpenHAB)

	applications := map[string]interface{}{
		"/echo/openhab": alexa.EchoApplication{
			AppID:    config.ApplicationId,
			OnIntent: OpenHABIntentHandler,
		},
	}

	alexa.Run(applications, "8080")
}

func OpenHABIntentHandler(echoReq *alexa.EchoRequest, echoResp *alexa.EchoResponse) {
	log.Printf("Handling %s intent with Slots %+v\n", echoReq.GetIntentName(), echoReq.AllSlots())
	switch echoReq.GetIntentName() {
	case "Switch":
		switchHandler(echoReq, echoResp)
	case "SetTemp":
		settempHandler(echoReq, echoResp)
	default:
		echoResp.OutputSpeech(lookupErrorResponse(UnknownCommand))
	}
}

func lookupAction(actionName string) string {

	if action, ok := config.Actions[actionName]; ok {
		return action
	}
	return actionName
}

func lookupItem(echoReq *alexa.EchoRequest) (string, error) {
	location, err := echoReq.GetSlotValue("Location")
	if err != nil {
		location = "default"
	}
	itemName, err := echoReq.GetSlotValue("ItemName")
	if err != nil {
		return "", ItemNotFound
	}

	if items, ok := config.Items[itemName]; ok {
		if item, ok := items[location]; ok {
			return item, nil
		}
	}
	return "", ItemNotFound
}

func lookupErrorResponse(key string) string {
	if response, ok := config.ErrorResponses[key]; ok {
		return response
	}
	return defaultErrorResponse
}

func lookupResponse(key string) string {
	if response, ok := config.SuccessResponse[key]; ok {
		return response
	}
	return defaultSuccessResponse
}
