package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageReplyTextOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11h9v2H9v-2m9-4H6v2h12V7m4-3v18l-4-4H4a2 2 0 0 1-2-2V4c0-1.1.9-2 2-2h16a2 2 0 0 1 2 2m-2 0H4v12h14.83L20 17.17V4Z"/>`),
		g.Group(children),
	)
}