package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionSharpTurn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 18v6.586L7.707 6.293A1 1 0 0 0 6 7v21h2V9.414L24.586 26H18v2h10V18Z"/>`),
		g.Group(children),
	)
}