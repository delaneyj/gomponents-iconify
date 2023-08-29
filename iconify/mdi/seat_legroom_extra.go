package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SeatLegroomExtra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 12V3H2v9a5 5 0 0 0 5 5h6v-2H7a3 3 0 0 1-3-3m18.83 5.24c-.38-.74-1.29-.97-2.03-.63l-1.09.5l-3.41-6.98A2.022 2.022 0 0 0 14.5 9H11V3H5v8a3 3 0 0 0 3 3h7l3.41 7l3.72-1.7c.77-.36 1.1-1.3.7-2.06Z"/>`),
		g.Group(children),
	)
}