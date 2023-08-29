package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h1v9h6v2.59l-4.21 4.2l1.42 1.42l2.79-2.8V22h2v-2.59l2.79 2.8l1.42-1.42l-4.21-4.2V14h6V5h1a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1H4Z"/>`),
		g.Group(children),
	)
}