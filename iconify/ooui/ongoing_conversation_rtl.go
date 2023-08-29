package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OngoingConversationRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M0 14c0 1.1.9 2 2 2h14l4 4V2c0-1.1-.9-2-2-2H2C.9 0 0 .9 0 2zm13.6-6.3c0-.8.6-1.4 1.4-1.4c.8 0 1.4.6 1.4 1.4s-.6 1.4-1.4 1.4c-.8-.1-1.4-.7-1.4-1.4zM9.9 9.1s-.1 0 0 0c-.8 0-1.4-.6-1.4-1.4c0-.8.6-1.4 1.4-1.4c.8 0 1.4.6 1.4 1.4s-.7 1.4-1.4 1.4zm-5.2 0c-.8 0-1.4-.6-1.4-1.4c0-.8.6-1.4 1.4-1.4c.8 0 1.4.6 1.4 1.4c0 .7-.7 1.4-1.4 1.4z"/>`),
		g.Group(children),
	)
}