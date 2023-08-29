package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveSelectionUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14V2h12v12H6Zm10 4v-2h2v2h-2ZM6 18v-2h2v2H6Zm10 4v-2h2v2h-2Zm-5 0v-2h2v2h-2Zm-5 0v-2h2v2H6Z"/>`),
		g.Group(children),
	)
}