package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeRoof(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 16h3L12 7L2 16h3l7-6.31L19 16M7 8.81V7H4v4.5l3-2.69Z"/>`),
		g.Group(children),
	)
}