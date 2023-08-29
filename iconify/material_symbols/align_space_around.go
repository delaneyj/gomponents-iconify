package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignSpaceAround(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 4V2h20v2H2Zm0 18v-2h20v2H2ZM7 9V6h10v3H7Zm0 9v-3h10v3H7Z"/>`),
		g.Group(children),
	)
}