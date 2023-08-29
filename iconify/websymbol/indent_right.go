package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IndentRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 118v213H331V118h669zm0 334v213H331V452h669zm0 335v213H331V787h669zM89 452l150 107L89 665V452z"/>`),
		g.Group(children),
	)
}