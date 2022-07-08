package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"jimcal.me/discord-go/x/mux"
)

// Router is registered as a global variable to allow easy access to the
// multiplexer throughout the bot.

// Session is declared in the global space so it can be easily used
// throughout this program.
// In this use case, there is no error that would be returned.
var Session, _ = discordgo.New("")
var Router = mux.New()

// Read in all configuration options from both environment variables and
// command line arguments.
func init() {
	Session.Token = "Bot $TOKEN"
	Session.AddHandler(Router.OnMessageCreate)

	// Register routes
	Router.Route("help", "Display this message.", Router.Help)
	Router.Route("ping", "Ping Adam", Router.Ping)
}

func main() {
	err := Session.Open()
	if err != nil {
		log.Printf("error opening connection to Discord, %s\n", err)
		os.Exit(1)
	}

	log.Printf(`Now running. Press CTRL-C to exit.`)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Clean up
	Session.Close()
}
