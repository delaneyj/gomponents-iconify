package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17 2v7a3 3 0 0 0 3 3h7v15a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3h9Zm2 .117V9a1 1 0 0 0 1 1h6.883a3 3 0 0 0-.762-1.293L20.293 2.88A3 3 0 0 0 19 2.117Z"/>`),
		g.Group(children),
	)
}