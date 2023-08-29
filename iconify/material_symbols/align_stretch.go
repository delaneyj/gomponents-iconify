package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignStretch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 10V4H2V2h20v2h-5v6H7ZM2 22v-2h5v-6h10v6h5v2H2Z"/>`),
		g.Group(children),
	)
}