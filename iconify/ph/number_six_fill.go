package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberSixFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 24H56a16 16 0 0 0-16 16v176a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V40a16 16 0 0 0-16-16Zm-72 168a44 44 0 0 1-38.3-65.62L123.38 68a8 8 0 0 1 13.86 8l-16.52 28.61A44.79 44.79 0 0 1 128 104a44 44 0 0 1 0 88Zm28-44a28 28 0 1 1-28-28a28 28 0 0 1 28 28Z"/>`),
		g.Group(children),
	)
}