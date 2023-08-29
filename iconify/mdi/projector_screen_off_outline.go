package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProjectorScreenOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.84 22.73l1.27-1.27L2.39 1.73L1.11 3L5 6.89V14h6v2.59l-4.21 4.2l1.42 1.42l2.79-2.8V22h2v-2.59l2.79 2.8l1.42-1.42l-4.21-4.2v-1.7l7.84 7.84M7 12V8.89L10.11 12H7m1.2-7l-3-3H20c.55 0 1 .45 1 1v1c0 .55-.45 1-1 1h-1v9h-1.8l-2-2H17V5H8.2Z"/>`),
		g.Group(children),
	)
}