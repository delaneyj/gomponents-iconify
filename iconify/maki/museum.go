package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Museum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 0L1 3.5V4h13v-.5L7.5 0zM2 5v5l-1 1.6V13h13v-1.4L13 10V5H2zm2 1h1v5.5H4V6zm3 0h1v5.5H7V6zm3 0h1v5.5h-1V6z"/>`),
		g.Group(children),
	)
}