package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var commands = []*discordgo.ApplicationCommand{
	{
		Name:        "ping",
		Description: "Responde con Pong y mide la latencia",
	},
	{
		Name:        "saludar",
		Description: "El bot te enviará un saludo personalizado",
	},
}

var commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"ping": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("🏓 ¡Pong! Latencia: `%v`", s.HeartbeatLatency()),
			},
		})
	},
	"saludar": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Obtener el usuario que ejecutó el comando (en servidor o mensaje directo)
		usuario := i.User
		if i.Member != nil {
			usuario = i.Member.User
		}

		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: fmt.Sprintf("¡Hola %s! 👋 ¿En qué te puedo ayudar hoy?", usuario.Mention()),
			},
		})
	},
}

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: no se encontró archivo .env, buscando en variables de entorno")
	}

	token := os.Getenv("TOKEN_BOT")
	if token == "" {
		log.Fatal("Error: TOKEN_BOT no está definido")
	}

	// Crear sesión
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creando sesión: %v", err)
	}

	dg.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if i.Type == discordgo.InteractionApplicationCommand {
			if handler, ok := commandHandlers[i.ApplicationCommandData().Name]; ok {
				handler(s, i)
			}
		}
	})

	if err := dg.Open(); err != nil {
		log.Fatalf("Error al conectar con Discord: %v", err)
	}
	defer dg.Close()

	// Registrar los comandos en Discord
	// NOTA: Dejar "" como segundo parámetro los registra de forma global.
	// (Los comandos globales pueden tardar hasta 1 hora en propagarse en Discord;
	// si pones el ID de tu servidor en lugar de "", se actualizan al instante).
	guildID := "" //*para pruebas insta

	log.Println("Registrando comandos...")
	registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
	for idx, cmd := range commands {
		createdCmd, err := dg.ApplicationCommandCreate(dg.State.User.ID, guildID, cmd)
		if err != nil {
			log.Fatalf("No se pudo crear el comando '%v': %v", cmd.Name, err)
		}
		registeredCommands[idx] = createdCmd
	}

	fmt.Println("🤖 Bot conectado y listo. Presiona CTRL+C para detener.")

	// Esperar señal de salida
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	// borrar comandos al apagar el bot
	fmt.Println("\nEliminando comandos registrados...")
	for _, cmd := range registeredCommands {
		err := dg.ApplicationCommandDelete(dg.State.User.ID, guildID, cmd.ID)
		if err != nil {
			log.Printf("No se pudo eliminar el comando '%v': %v", cmd.Name, err)
		}
	}

	fmt.Println("Bot desconectado correctamente.")
}
