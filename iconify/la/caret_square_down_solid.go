package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CaretSquareDownSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm3.719 5.781L9.28 14.22l6 6l.719.687l.719-.687l6-6l-1.438-1.438l-5.28 5.28z"/>`),
		g.Group(children),
	)
}