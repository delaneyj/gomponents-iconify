package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignSpaceBetween(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 7V4H2V2h20v2h-5v3H7ZM2 22v-2h5v-3h10v3h5v2H2Z"/>`),
		g.Group(children),
	)
}