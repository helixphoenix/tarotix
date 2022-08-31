package controllers

import (
	"log"
	"time"
	"github.com/tarotix/views"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
	"math/rand"
)


type CommandController struct {
	EventHandler *socketmode.SocketmodeHandler
}

func NewCommandController(eventhandler *socketmode.SocketmodeHandler) CommandController {
	command := CommandController{
		EventHandler: eventhandler,
	}

	command.EventHandler.HandleSlashCommand("/daily", command.launchDailyTarot
 )

    command.EventHandler.HandleInteractionBlockAction(views.DailyTarotActionID, command.launchDaily,
	)

	return command
}



func (command CommandController) launchDailyTarot (evt *socketmode.Event, clt *socketmode.Client) {
	command, ok := evt.Data.(slack.SlashCommand)

	if ok != true {
		log.Printf("Error command is not working ", ok)
	}

	clt.Ack(*evt.Request)
   
    daily := 

}


func pickNCard(n int ){
	min := 0
	max := 52

	cards := []
    for card := range n {
		card := rand.Intn(max-min)
		cards = append(cards,card)
	}
    
}