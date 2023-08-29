package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackTie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v22h22V5zm2 2h18v18H7zm5 2l2.813 3.625L12 20.375L16 24l4-3.625l-2.813-7.75L20 9z"/>`),
		g.Group(children),
	)
}