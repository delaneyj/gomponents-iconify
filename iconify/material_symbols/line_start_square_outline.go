package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStartSquareOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 15.5h7v-7H4v7Zm9 2H2v-11h11V11h9v2h-9v4.5ZM7.5 12Z"/>`),
		g.Group(children),
	)
}