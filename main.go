package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"bot_barionyx/commands"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/voice"
	"github.com/joho/godotenv"
	"github.com/thomas-vilte/dave-go/session"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: no se encontró archivo .env")
	}

	token := os.Getenv("TOKEN_BOT")
	if token == "" {
		log.Fatal("TOKEN_BOT no está definido")
	}

	// 1. Configurar y crear cliente Disgo
	client, err := disgo.New(token,
		bot.WithCacheConfigOpts(cache.WithCaches(cache.FlagVoiceStates)),
		bot.WithVoiceManagerConfigOpts(
			voice.WithDaveSessionCreateFunc(session.CreateFunc()),
		),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
				gateway.IntentGuildVoiceStates, // Requerido para saber en qué canal de voz está el usuario
			),
		),
		bot.WithEventListeners(&events.ListenerAdapter{
			OnApplicationCommandInteraction: commands.HandleInteraction,
		}),
	)
	if err != nil {
		log.Fatal("Error creando cliente:", err)
	}

	// 2. Conectar a Discord
	if err = client.OpenGateway(context.TODO()); err != nil {
		log.Fatal("Error conectando al Gateway:", err)
	}

	// 3. Registrar comandos
	log.Println("Registrando comandos globales...")
	_, err = client.Rest.SetGlobalCommands(client.ApplicationID, commands.CommandList)
	if err != nil {
		log.Fatal("Error registrando comandos:", err)
	}

	fmt.Println("🤖 Bot conectado y listo con Disgo. Presiona CTRL+C para detener.")
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s

	fmt.Println("\nDesconectando...")
	client.Close(context.TODO())
}
