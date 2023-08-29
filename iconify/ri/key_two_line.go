package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.758 11.828l7.849-7.849l1.414 1.414l-1.414 1.414l2.474 2.475l-1.414 1.415l-2.475-2.475l-1.414 1.414l2.121 2.121l-1.414 1.414l-2.121-2.12l-2.192 2.191a5.002 5.002 0 0 1-7.708 6.293a5 5 0 0 1 6.294-7.707Zm-.637 6.293A3 3 0 1 0 5.88 13.88a3 3 0 0 0 4.242 4.242Z"/>`),
		g.Group(children),
	)
}