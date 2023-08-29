package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairAltSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21v-9h3v-2H5V3h14v7h-3v2h3v9h-2v-3H7v3H5Zm5-9h4v-2h-4v2Z"/>`),
		g.Group(children),
	)
}