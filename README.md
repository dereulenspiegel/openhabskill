# openhabskill

This very creatively named repository contains a skill for Amazon Alexa which allows
to control your Home through [openHAB](https://github.com/openhab/openhab).

# Requirements

* An Amazon developer account
* A machine in your local network to run the skillserver (preferably via Docker)
* The possibility to expose the URL of the skillserver via HTTPS to the outside world
* A valid TLS certificate
* A valid DNS name to reach the host the skill runs on
* Some time on your hands

# Features

[x] Switch items on and off
[x] Set a presence switch via pre determined greetings
[x] Set temperature
[ ] Template voice responses
[ ] Get state of items
[ ] Set scenes
[ ] Generate utterances from templates

# Language support
[x] German
[ ] UK English
[ ] US English

# Howto

## Skillserver

### Docker

* Pull the image dereulenspiegel/openhabskill:0.1-arm
* Mount a host folder to /configuration in the container
* Expose port 8080
* create a config.yml in the folder mounted under /configuration

### Classic

* Have a valid go environment on your machine
* `go get -u github.com/dereulenspiegel/openhabskill`
* Run the skillserver on your machine

* Have a look at the config.example.yml in the configuration folder of this repository to create your own configuration
* App and User ID are shown in the log output of the skill server on the first requests.
* Configure Nginx/Apache/etc. so that it can act as a TLS terminating reverse proxy in front of your skillserver
* You need to proxy to `/echo/openhabskill`

## Amazon Developer account
* Set up an [Amazon developer account](https://developer.amazon.com)
* Under Alexa create Add a new Skill with a custom interaction model
* Set the desired language
* Set a name you would recognize
* Choose an invocation name (what you will be saying)
* Click next to interaction model
* Copy and paste the intent_schema.json from the configuration directory into the Intent Schema field
* Create the custom slot types as specified in slot_types.de (or in the future slot_types.<us|uk>)
* Copy and paste the content of utterances.de into the Sample Utterances field.
  * Please note that the Amazon Alexa Skill webui is sometimes buggy and won't let you save the interaction model
    until all custom slot types are defined. A workaround is to create a nearly empty intent schema wich has
    a single intent, without any slots. This lets you create all custom slot types. Then you can paste
    the intent schema
* Next specify the HTTPS endpoint where Amazon Alexa can reach your skill
* Next specify the exakt kind of SSL certificate you are using
* After this you can start testing your skill