package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

type Qt struct {
	Content string
	//	Character string
}

func main() {
	token := os.Getenv("TOKEN")
	b, err := gotgbot.NewBot(token, &gotgbot.BotOpts{
		Client:      http.Client{},
		GetTimeout:  gotgbot.DefaultGetTimeout,
		PostTimeout: gotgbot.DefaultPostTimeout,
	})
	if err != nil {
		panic("failed to create new bot: " + err.Error())
	}

	updater := ext.NewUpdater(nil)
	dispatcher := updater.Dispatcher

	dispatcher.AddHandler(handlers.NewCommand("quote", quotex))
	dispatcher.AddHandler(handlers.NewCommand("start", start))
	dispatcher.AddHandler(handlers.NewCommand("repo", repo))

	err = updater.StartPolling(b, &ext.PollingOpts{DropPendingUpdates: true})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	fmt.Printf("%s has been started...\n", b.User.Username)

	updater.Idle()
}

func quotex(b *gotgbot.Bot, ctx *ext.Context) error {
	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		fmt.Println(err)
	}
	data, _ := ioutil.ReadAll(resp.Body)
	var quote Qt
	json.Unmarshal([]byte(data), &quote)

	_, errr := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("<b>%v</b>", quote.Content), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})

	if errr != nil {
		fmt.Print(errr)
	}
	return nil
}

func start(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("Hello"), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})

	if err != nil {
		panic(err)
	}

	return nil
}

func repo(b *gotgbot.Bot, ctx *ext.Context) error {
	_, err := ctx.EffectiveMessage.Reply(b, fmt.Sprintf("https://github.com/Nksama/AnimeQuotes"), &gotgbot.SendMessageOpts{
		ParseMode: "html",
	})

	if err != nil {
		panic(err)
	}
	return nil
}

