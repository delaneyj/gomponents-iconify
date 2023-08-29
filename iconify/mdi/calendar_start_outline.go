package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarStartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 18l-4-4v3H4v-3H2v8h2v-3h5v3l4-4m6-15h-1V1h-2v2H8V1H6v2H5c-1.1 0-2 .9-2 2v7h2V9h14v10h-4.2l-2 2H19c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2M5 7V5h14v2H5Z"/>`),
		g.Group(children),
	)
}