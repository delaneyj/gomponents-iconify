package uiw

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 0c5.523 0 10 4.477 10 10s-4.477 10-10 10S0 15.523 0 10S4.477 0 10 0Zm-.48 7.636l-3.896 3.87a.682.682 0 0 0 .962.968l3.419-3.398l3.49 3.402a.682.682 0 0 0 .952-.976l-3.971-3.87a.682.682 0 0 0-.957.004Z"/>`),
		g.Group(children),
	)
}