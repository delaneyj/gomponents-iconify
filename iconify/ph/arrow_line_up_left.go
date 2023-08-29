package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 216a8 8 0 0 1-8 8H40a8 8 0 0 1 0-16h176a8 8 0 0 1 8 8ZM64 160a8 8 0 0 0 8-8V75.31l98.34 98.35a8 8 0 0 0 11.32-11.32L83.31 64H160a8 8 0 0 0 0-16H64a8 8 0 0 0-8 8v96a8 8 0 0 0 8 8Z"/>`),
		g.Group(children),
	)
}