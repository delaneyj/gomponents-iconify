package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreaterThan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.5 4.14l-1 1.72L15 12L4.5 18.14l1 1.72L19 12L5.5 4.14Z"/>`),
		g.Group(children),
	)
}