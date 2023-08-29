package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 216a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h176a8 8 0 0 1 8 8ZM80 176a8 8 0 0 0 5.66-2.34L184 75.31V152a8 8 0 0 0 16 0V56a8 8 0 0 0-8-8H96a8 8 0 0 0 0 16h76.69l-98.35 98.34A8 8 0 0 0 80 176Z"/>`),
		g.Group(children),
	)
}