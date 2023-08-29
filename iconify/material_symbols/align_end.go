package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignEnd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-2h20v2H2Zm5-11V8h10v3H7Zm0 6v-3h10v3H7Z"/>`),
		g.Group(children),
	)
}