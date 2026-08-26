package commands

import (
	"bot_barionyx/music"
	"fmt"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var PlayCommand = discord.SlashCommandCreate{
	Name:        "play",
	Description: "Reproduce una canción desde YouTube",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:        "url",
			Description: "El enlace de YouTube",
			Required:    true,
		},
	},
}

func PlayHandler(client *bot.Client, e *events.ApplicationCommandInteractionCreate) {
	url := e.SlashCommandInteractionData().String("url")
	guildID := *e.GuildID()
	userID := e.User().ID

	// Disgo guarda en caché dónde está el usuario
	voiceState, ok := client.Caches.VoiceState(guildID, userID)
	if !ok || voiceState.ChannelID == nil {
		e.CreateMessage(discord.MessageCreate{Content: "❌ ¡Debes estar en un canal de voz primero!"})
		return
	}
	channelID := *voiceState.ChannelID

	err := music.UnirseCanal(client, guildID, channelID)
	if err != nil {
		e.CreateMessage(discord.MessageCreate{Content: "❌ Error al unirse al canal de voz."})
		return
	}

	e.CreateMessage(discord.MessageCreate{Content: fmt.Sprintf("🎵 Preparando y reproduciendo: %s", url)})

	// Reproducir en segundo plano
	go func() {
		err := music.PlayYouTube(client, guildID, url)
		if err != nil {
			fmt.Printf("Error de reproducción: %v\n", err)
		}
	}()
}

var LeaveCommand = discord.SlashCommandCreate{
	Name:        "leave",
	Description: "Desconecta al bot del canal de voz",
}

func LeaveHandler(client *bot.Client, e *events.ApplicationCommandInteractionCreate) {
	err := music.Desconectar(client, *e.GuildID())
	if err != nil {
		e.CreateMessage(discord.MessageCreate{Content: "❌ No estoy en un canal de voz."})
		return
	}
	e.CreateMessage(discord.MessageCreate{Content: "👋 Me he desconectado del canal de voz."})
}
