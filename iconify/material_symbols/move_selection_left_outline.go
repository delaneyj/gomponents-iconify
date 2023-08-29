package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveSelectionLeftOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 18V6h12v12H2Zm2-2h8V8H4v8Zm12-8V6h2v2h-2Zm0 10v-2h2v2h-2Zm4-10V6h2v2h-2Zm0 5v-2h2v2h-2Zm0 5v-2h2v2h-2ZM8 12Z"/>`),
		g.Group(children),
	)
}