package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22V2h2v20H2Zm4-5v-3h10v3H6Zm0-7V7h16v3H6Z"/>`),
		g.Group(children),
	)
}