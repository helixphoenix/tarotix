package main

import (
	"fmt"
	"log"
	"os"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tarotix/drivers"
	"github.com/tarotix/controllers"



)



func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("err loading env file")
	}
    tarotix, err := drivers.ConnectSlackSocket()
	if err != nil {
		log.Error("error connecting").Str("error", err.Error()).Msg("Unable to connect to slack")


		os.Exit(1)
	}
    socketmodeHandler := socketmode.NewSocketmodeHandler(tarotix)

	controllers.NewGreetingController(socketmodeHandler)
	controllers.NewCommandController(socketmodeHandler)
	socketmodeHandler.RunEventLoop()

	

}
