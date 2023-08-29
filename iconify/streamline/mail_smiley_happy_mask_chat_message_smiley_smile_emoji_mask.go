package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSmileyHappyMaskChatMessageSmileySmileEmojiMask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 8s.5 1.5 3 1.5S10 8 10 8"/><path d="M13.5 7a6.5 6.5 0 0 1-13 0V.5h13Z"/><path d="M5 5a1 1 0 0 0-2 0m8 0a1 1 0 0 0-2 0"/></g>`),
		g.Group(children),
	)
}