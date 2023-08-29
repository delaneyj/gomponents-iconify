package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRotateClockwiseSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 3a5 5 0 0 0-5 5a.5.5 0 0 1-1 0a6 6 0 1 1 9.969 4.5H13a.5.5 0 0 1 0 1h-2.5a.5.5 0 0 1-.5-.5v-2.5a.5.5 0 0 1 1 0V12a5 5 0 0 0-3-9ZM6 8a2 2 0 1 1 4 0a2 2 0 0 1-4 0Zm2-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/>`),
		g.Group(children),
	)
}