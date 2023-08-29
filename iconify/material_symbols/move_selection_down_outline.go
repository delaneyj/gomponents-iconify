package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveSelectionDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22V10h12v12H6Zm2-2h8v-8H8v8ZM6 8V6h2v2H6Zm10 0V6h2v2h-2ZM6 4V2h2v2H6Zm5 0V2h2v2h-2Zm5 0V2h2v2h-2Zm-4 12Z"/>`),
		g.Group(children),
	)
}