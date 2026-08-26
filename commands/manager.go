package commands

import (
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var CommandList = []discord.ApplicationCommandCreate{
	PingCommand,
	SaludarCommand,
	DadoCommand,
	PlayCommand,
	LeaveCommand,
}

var CommandHandlers = map[string]func(client *bot.Client, e *events.ApplicationCommandInteractionCreate){
	"ping":    PingHandler,
	"saludar": SaludarHandler,
	"dado":    DadoHandler,
	"play":    PlayHandler,
	"leave":   LeaveHandler,
}

// HandleInteraction enruta el comando que ingresó el usuario a su función
func HandleInteraction(e *events.ApplicationCommandInteractionCreate) {
	if handler, ok := CommandHandlers[e.Data.CommandName()]; ok {
		handler(e.Client(), e)
	}
}
