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
	guildIDPtr := e.GuildID()
	if guildIDPtr == nil {
		e.CreateMessage(discord.MessageCreate{Content: "❌ Este comando solo puede usarse dentro de un servidor."})
		return
	}
	guildID := *guildIDPtr
	userID := e.User().ID

	// Disgo guarda en caché dónde está el usuario
	voiceState, ok := client.Caches.VoiceState(guildID, userID)
	if !ok || voiceState.ChannelID == nil {
		e.CreateMessage(discord.MessageCreate{Content: "❌ ¡Debes estar en un canal de voz primero!"})
		return
	}
	channelID := *voiceState.ChannelID

	if err := e.DeferCreateMessage(false); err != nil {
		fmt.Printf("Error respondiendo la interacción /play: %v\n", err)
		return
	}

	err := music.UnirseCanal(client, guildID, channelID)
	if err != nil {
		message := "❌ Error al unirse al canal de voz: " + err.Error()
		_, updateErr := client.Rest.UpdateInteractionResponse(
			e.ApplicationID(),
			e.Token(),
			discord.MessageUpdate{Content: &message},
		)
		if updateErr != nil {
			fmt.Printf("Error actualizando la respuesta de /play: %v\n", updateErr)
		}
		return
	}

	message := fmt.Sprintf("🎵 Preparando y reproduciendo: %s", url)
	_, err = client.Rest.UpdateInteractionResponse(
		e.ApplicationID(),
		e.Token(),
		discord.MessageUpdate{Content: &message},
	)
	if err != nil {
		fmt.Printf("Error actualizando la respuesta de /play: %v\n", err)
	}

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
