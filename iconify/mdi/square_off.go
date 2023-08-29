package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.84 22.73L19.11 21H3V4.89L1.11 3l1.28-1.27l19.72 19.73l-1.27 1.27M21 3H6.2L21 17.8V3Z"/>`),
		g.Group(children),
	)
}