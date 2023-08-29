package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchoolEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M8.5 9V8H10V6H7.5V5H10V3H8.5V2H10V1H6v9h4V9H8.5zM4 7H1V1h3v6zm0 1l-1.5 2L1 8h3z" fill="currentColor"/>`),
		g.Group(children),
	)
}