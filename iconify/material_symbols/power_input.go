package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerInput(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 15v-2h5v2H2Zm7 0v-2h5v2H9Zm7 0v-2h5v2h-5ZM2 11V9h19v2H2Z"/>`),
		g.Group(children),
	)
}