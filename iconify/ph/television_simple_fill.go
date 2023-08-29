package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 80v120a16 16 0 0 1-16 16H40a16 16 0 0 1-16-16V80a16 16 0 0 1 16-16h68.69L74.34 29.66a8 8 0 0 1 11.32-11.32L128 60.69l42.34-42.35a8 8 0 1 1 11.32 11.32L147.31 64H216a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}