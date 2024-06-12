package telegram

import (
	"Telegramm-botTZ/lib/e"
	"Telegramm-botTZ/storage"
	"log"
	"net/url"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doCmd(text string, chatID int, username string) error {
	log.Printf("got new command '%s' for '%s", text, username)
	text = strings.TrimSpace(text)

	if isAddCmd(text) {
		switch text {
		case RndCmd:

		case HelpCmd:
		case StartCmd:
		default:

		}

	}

}

func (p *Processor) savePage(chatID int, text string, username string) (err error) {
	defer func() { err = e.Warp("can't do command: save", err) }()

	page := &storage.Page{
		URL:      pageURL,
		UserName: "",
	}

	isExists, err := p.storage.IsExist(page)
	if err != nil {
		return err
	}
	if isExists {
		return p.tg.SendMessage(chatID, "")
	}

} // stopping place bot response texts

func isAddCmd(text string) bool {
	return isURL(text)
}

func isURL(text string) bool {
	u, err := url.Parse(text)

	return err == nil && u.Host != ""

}
