package streamline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailSmileyStraightFaceChatMessageIndifferentSmileyEmojiFacePoker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 14 14"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3.5 9.5h7m-3.5 4a6.5 6.5 0 1 0 0-13a6.5 6.5 0 0 0 0 13Z"/><path d="M4.8 5.45a.25.25 0 0 1 0-.5m0 .5a.25.25 0 0 0 0-.5m4.4.5a.25.25 0 0 1 0-.5m0 .5a.25.25 0 0 0 0-.5"/></g>`),
		g.Group(children),
	)
}