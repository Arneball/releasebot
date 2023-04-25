package botstuff

import (
	"fmt"
	goteamsnotify "github.com/atc0005/go-teams-notify/v2"
	"github.com/atc0005/go-teams-notify/v2/adaptivecard"
	"github.com/spf13/cobra"
	"io"
	"net/url"
	"os"
)

func sendMessage(theUrl, theTitle, theWebhook, theMessage string) error {
	c := goteamsnotify.NewTeamsClient()
	card := adaptivecard.NewCard()
	openURL, err := adaptivecard.NewActionOpenURL(theUrl, theTitle)
	if err != nil {
		return err
	}
	img := adaptivecard.Element{
		Type: "Image",
		URL:  fmt.Sprintf("https://qr.bache.se?url=%s", url.QueryEscape(theUrl)),
	}
	err = card.AddElement(false, img)
	if err != nil {
		return err
	}
	err = card.AddElement(true, adaptivecard.NewTextBlock(theMessage, true))
	if err := card.AddAction(false, openURL); err != nil {
		return err
	}
	msg, err := adaptivecard.NewMessageFromCard(card)
	if err != nil {
		return err
	}
	if err := c.Send(theWebhook, msg); err != nil {
		return err
	}
	return nil
}

var TeamsCommand = &cobra.Command{
	Use:   "teams-notify",
	Short: "Notify teams chat",
	RunE: func(cmd *cobra.Command, args []string) error {
		flags := cmd.Flags()
		title, err := flags.GetString("title")
		if err != nil {
			return err
		}
		url, err := flags.GetString("url")
		if err != nil {
			return err
		}
		hook, err := flags.GetString("webhook_url")
		if err != nil {
			return err
		}
		theMessage, err := io.ReadAll(cmd.InOrStdin())
		if err != nil {
			return err
		}
		message := string(theMessage)
		_, _ = fmt.Fprintf(os.Stderr, "title=%s, url=%s, msg=%s", title, url, message)
		return sendMessage(url, title, hook, message)
	},
}

func init() {
	TeamsCommand.Flags().String("title", "", "Title of teams notification message")
	TeamsCommand.Flags().String("url", "", "Url to where the binary lives")
	TeamsCommand.Flags().String("webhook_url", "", "Url to the teams webhook")
	for _, flag := range []string{"title", "url", "webhook_url"} {
		if err := TeamsCommand.MarkFlagRequired(flag); err != nil {
			panic(err)
		}
	}
}
