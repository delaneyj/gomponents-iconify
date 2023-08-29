package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarBlankMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 17V8H7v9h14m0-14c1.1 0 2 .9 2 2v12c0 1.1-.9 2-2 2H7a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h1V1h2v2h8V1h2v2h1M3 21h14v2H3a2 2 0 0 1-2-2V9h2v12Z"/>`),
		g.Group(children),
	)
}