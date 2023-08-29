package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourGMobiledataSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17v-3H3V7h2v5h2V7h2v5h2v2H9v3H7Zm5 0V7h9v2h-7v6h5v-2h-2.5v-2H21v6h-9Z"/>`),
		g.Group(children),
	)
}