package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSmileySadFaceChatMessageSmileyEmojiSadFaceUnsatisfied(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 13.5a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z"/><path d="M3.7 10.5C4.2 8.7 6.1 7.6 8 8.1c1.1.3 2 1.2 2.4 2.4M4.85 5.45a.25.25 0 0 1 0-.5m0 .5a.25.25 0 0 0 0-.5m4.4.5a.25.25 0 0 1 0-.5m0 .5a.25.25 0 0 0 0-.5"/></g>`),
		g.Group(children),
	)
}