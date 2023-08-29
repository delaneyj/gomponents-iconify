package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletDashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 18H5V6h14m2-2H3c-1.11 0-2 .89-2 2v12a2 2 0 0 0 2 2h18a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2M7 8h6v5H7V8m7 0h3v2h-3V8m3 3v5h-3v-5h3M7 14h6v2H7v-2Z"/>`),
		g.Group(children),
	)
}