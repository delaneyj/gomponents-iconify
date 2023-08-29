package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GarageAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20h-2v-9H5v9H3V9l7-4l7 4v11M6 12h8v2H6v-2m0 3h8v2H6v-2m13 0v-5h2v5h-2m0 4v-2h2v2h-2Z"/>`),
		g.Group(children),
	)
}