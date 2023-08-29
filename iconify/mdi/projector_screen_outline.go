package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorScreenOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-.55 0-1 .45-1 1v1c0 .55.45 1 1 1h1v9h6v2.59l-4.21 4.2l1.42 1.42l2.79-2.8V22h2v-2.59l2.79 2.8l1.42-1.42l-4.21-4.2V14h6V5h1c.55 0 1-.45 1-1V3c0-.55-.45-1-1-1m-3 10H7V5h10v7Z"/>`),
		g.Group(children),
	)
}