package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Select(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 9.657l1.414 1.414l4.243-4.243l4.242 4.243l1.415-1.414L11.657 4L6 9.657Zm0 4.786l1.414-1.414l4.243 4.243l4.242-4.243l1.415 1.414l-5.657 5.657L6 14.443Z"/>`),
		g.Group(children),
	)
}