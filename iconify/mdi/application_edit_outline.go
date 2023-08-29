package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApplicationEditOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 20v2H3c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h18c1.1 0 2 .9 2 2v8.1l-.2-.2c-.5-.5-1.1-.8-1.8-.8V6H3v14h8m10.4-6.7l1.3 1.3c.2.2.2.6 0 .8l-1 1l-2.1-2.1l1-1c.1-.1.2-.2.4-.2s.3.1.4.2m-.3 3.6l-6 6.1H13v-2.1l6.1-6.1l2 2.1Z"/>`),
		g.Group(children),
	)
}