package music

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/voice"
	"github.com/disgoorg/snowflake/v2"
	"github.com/jonas747/dca"
)

func UnirseCanal(client *bot.Client, guildID, channelID snowflake.ID) error {

	conn := client.VoiceManager.CreateConn(guildID)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return conn.Open(ctx, channelID, false, true)
}

func Desconectar(client *bot.Client, guildID snowflake.ID) error {
	conn := client.VoiceManager.GetConn(guildID)
	if conn == nil {
		return fmt.Errorf("no hay conexión en este servidor")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn.Close(ctx)
	client.VoiceManager.RemoveConn(guildID)
	return nil
}

func GetStreamURL(youtubeURL string) (string, error) {
	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "-g", youtubeURL)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// dcaProvider sirve de "puente" entre DCA y Disgo
type dcaProvider struct {
	session *dca.EncodeSession
}

func (p *dcaProvider) ProvideOpusFrame() ([]byte, error) {
	return p.session.OpusFrame() // Le pasamos a Disgo los cuadros exactos de audio
}

func (p *dcaProvider) Close() {
	p.session.Cleanup()
}

func PlayYouTube(client *bot.Client, guildID snowflake.ID, youtubeURL string) error {
	conn := client.VoiceManager.GetConn(guildID)
	if conn == nil {
		return fmt.Errorf("no hay conexión")
	}

	streamURL, err := GetStreamURL(youtubeURL)
	if err != nil {
		return err
	}

	options := dca.StdEncodeOptions
	options.RawOutput = true
	options.Bitrate = 96
	options.Application = "lowdelay"

	session, err := dca.EncodeFile(streamURL, options)
	if err != nil {
		return err
	}

	// Decirle a Discord que se "encienda" el arillo verde de hablar
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn.SetSpeaking(ctx, voice.SpeakingFlagMicrophone)

	// Disgo se encarga de pedirle audio a este proveedor cada 20 milisegundos!
	conn.SetOpusFrameProvider(&dcaProvider{session: session})

	return nil
}
