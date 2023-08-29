package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSmileySadMaskChatMessageSmileyEmojiSadMaskCry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.5 7a6.5 6.5 0 0 1-13 0V.5h13Z"/><path d="M4 9.5S4.5 8 7 8s3 1.5 3 1.5m-5-5a1 1 0 0 1-2 0m8 0a1 1 0 0 1-2 0"/></g>`),
		g.Group(children),
	)
}