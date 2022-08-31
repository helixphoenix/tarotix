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
   
    daily := pickNCard(1)
	msg := tarotMeanings[daily]

	tarotix := clt.GetApiClient()
	_, _, err := tarotix.PostMessage(
		command.ChannelID, 
		slack.MsgOptionText("You have very bright future"),

	)

	if err != nil {
		log.Printf("Error while sending message", err)
	}


}


func pickNCard(n int ){
	min := 0
	max := 52

	var cards []int
    for i := 1; i< n+1; i++ {
		card := rand.Intn(max-min)
		if card not in cards {
			cards = append(cards,card)
		}else{
			card := rand.Intn(max-min)
			cards = append(cards,card)

		}

	}
    return cards
}