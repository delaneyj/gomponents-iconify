package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M189.66 122.34a8 8 0 0 1 0 11.32l-72 72a8 8 0 0 1-11.32-11.32L164.69 136H32a8 8 0 0 1 0-16h132.69l-58.35-58.34a8 8 0 0 1 11.32-11.32ZM216 32a8 8 0 0 0-8 8v176a8 8 0 0 0 16 0V40a8 8 0 0 0-8-8Z"/>`),
		g.Group(children),
	)
}