package main

import (
	"log"
	"net/http"
	"time"
        "os/exec"
        "bytes"
	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	menu = &tb.ReplyMarkup{}
)


func Shellout(command string) (error, string) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command("bash", "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return err, stdout.String()
}

func main() {
	b, err := tb.NewBot(tb.Settings{
		URL:       "",
		Token:     "5050904599:AAG-YrM6KN4EJJx8peQOn901qHhLCkFo5QA",
		Updates:   0,
		Poller:    &tb.LongPoller{Timeout: 10 * time.Second},
		ParseMode: "HTML",
		Reporter: func(error) {
		},
		Client: &http.Client{},
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	b.Handle("/info", func(m *tb.Message) {
		if string(m.Payload) == string("") {
			b.Send(m.Sender, "Mieko")
			return
		}
		if !m.IsReply() {
			b.Send(m.Sender, m.Sender.FirstName)
		}
	})

	b.Handle("/start", func(m *tb.Message) {
		if m.Private() {
			menu.Inline(
				menu.Row(menu.URL("Support", "t.me/roseloverx_support"), menu.URL("Updates", "t.me/roseloverx_support")),
				menu.Row(menu.Data("Commands", "help_menu")),
				menu.Row(menu.URL("Add me to your group", "https://t.me/Yoko_Robot?startgroup=true")))
			b.Send(m.Sender, "Hey there! I am <b>Yoko</b>.\nIm an Anime themed Group Management Bot, feel free to add me to your groups!", menu)
			return
		}
		b.Reply(m, "Hey I'm Alive.")
	})

        b.Handle("/sh", func(m *tb.Message) {
                if string(m.Payload) == string("") {
                   b.Reply(m, "Give some cmd to Execute!")
                   return
                  }
                err, out := Shellout(m.Payload)
                if string(err.Error()) == string("") {
                  b.Reply(m, "Go#~: " + m.Payload + "\n" + string(err.Error()))
                  }
                b.Reply(m, "Go#~: " + m.Payload + "\n" + string(out))
        })
	b.Start()
}
