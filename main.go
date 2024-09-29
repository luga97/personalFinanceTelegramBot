package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	tb "gopkg.in/telebot.v3"
)

var (
	CommandTest1 = tb.Command{Text: "test1", Description: "test1"}
	CommandTest2 = tb.Command{Text: "test2", Description: "test2"}
	CommandTest3 = tb.Command{Text: "test3", Description: "test3"}
	CommandTest4 = tb.Command{Text: "test4", Description: "test4"}
	CommandTest5 = tb.Command{Text: "test5", Description: "test5"}
	CommandTest6 = tb.Command{Text: "test6", Description: "test6"}
	commands     = []tb.Command{CommandTest1, CommandTest2, CommandTest3, CommandTest4, CommandTest5, CommandTest6}
)

func main() {
	godotenv.Load()
	//fmt.Println("token=", os.Getenv("BOT_TOKEN"))
	b, err := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		//URL: "http://195.129.111.17:8012",

		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	var (
		// Universal markup builders.
		menu     = &tb.ReplyMarkup{ResizeKeyboard: true}
		selector = &tb.ReplyMarkup{}

		// Reply buttons.
		btnHelp     = menu.Text("ℹ Help")
		btnSettings = menu.Text("⚙ Settings")

		// Inline buttons.
		//
		// Pressing it will cause the client to
		// send the bot a callback.
		//
		// Make sure Unique stays unique as per button kind,
		// as it has to be for callback routing to work.
		//
		btnPrev = selector.Data("⬅", "prev", "data prev")
		btnNext = selector.Data("➡", "next", "data next")
	)

	menu.Reply(
		menu.Row(btnHelp),
		menu.Row(btnSettings),
	)
	selector.Inline(
		selector.Row(btnPrev, btnNext),
	)

	// Command: /start <PAYLOAD>
	b.Handle("/start", func(c tb.Context) error {
		return c.Send("Hello world")
	})

	b.Handle("/"+CommandTest1.Text, func(c tb.Context) error {

		return c.Send(CommandTest1.Description, menu)
	})
	b.Handle("/"+CommandTest2.Text, func(c tb.Context) error {
		return c.Send(CommandTest2.Description)
	})
	b.Handle("/"+CommandTest3.Text, func(c tb.Context) error {
		return c.Send(CommandTest3.Description)
	})
	b.Handle("/"+CommandTest4.Text, func(c tb.Context) error {
		return c.Send(CommandTest4.Description)
	})
	b.Handle("/"+CommandTest5.Text, func(c tb.Context) error {
		return c.Send(CommandTest5.Description)
	})
	b.Handle("/"+CommandTest6.Text, func(c tb.Context) error {
		return c.Send(CommandTest6.Description)
	})
	b.Handle(tb.OnText, func(c tb.Context) error {
		return c.Send("texto simple: " + c.Text())
	})

	fmt.Println("bot runing")
	b.SetCommands(commands)
	b.Start()
}
