package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretSquareUpSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm9 4.094l-.719.687l-6 6l1.438 1.438L16 13.937l5.281 5.282l1.438-1.438l-6-6z"/>`),
		g.Group(children),
	)
}