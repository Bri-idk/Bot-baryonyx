package commands

import (
	"fmt"
	"math/rand"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var PingCommand = discord.SlashCommandCreate{
	Name:        "ping",
	Description: "Responde con Pong",
}

func PingHandler(client *bot.Client, e *events.ApplicationCommandInteractionCreate) {
	e.CreateMessage(discord.MessageCreate{
		Content: "🏓 ¡Pong!",
	})
}

var SaludarCommand = discord.SlashCommandCreate{
	Name:        "saludar",
	Description: "El bot te enviará un saludo personalizado",
}

func SaludarHandler(client *bot.Client, e *events.ApplicationCommandInteractionCreate) {
	e.CreateMessage(discord.MessageCreate{
		Content: fmt.Sprintf("¡Hola %s! 👋 ¿En qué te puedo ayudar?", e.User().Mention()),
	})
}

var DadoCommand = discord.SlashCommandCreate{
	Name:        "dado",
	Description: "Lanza un dado del 1 al 6",
}

func DadoHandler(client *bot.Client, e *events.ApplicationCommandInteractionCreate) {
	resultado := rand.Intn(6) + 1
	e.CreateMessage(discord.MessageCreate{
		Content: fmt.Sprintf("🎲 Has lanzado un dado y obtuviste un **%d**", resultado),
	})
}
