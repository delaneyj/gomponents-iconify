package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CarBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 3v3H1v14h22V6h-3V3h-6v3h-4V3H4M3 8h18v10H3V8m12 2v2h-2v2h2v2h2v-2h2v-2h-2v-2h-2M5 12v2h6v-2H5Z"/>`),
		g.Group(children),
	)
}