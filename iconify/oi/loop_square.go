package oi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoopSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 8 8"),
		g.Raw(`<path fill="currentColor" d="M1 0v2h1V1h4v2H5l1.5 2.5L8 3H7V0H1zm.5 2.5L0 5h1v3h6V6H6v1H2V5h1L1.5 2.5z"/>`),
		g.Group(children),
	)
}