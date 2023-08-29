package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveSelectionRightOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 18V6h12v12H10Zm2-2h8V8h-8v8Zm-6 2v-2h2v2H6ZM6 8V6h2v2H6ZM2 18v-2h2v2H2Zm0-5v-2h2v2H2Zm0-5V6h2v2H2Zm14 4Z"/>`),
		g.Group(children),
	)
}