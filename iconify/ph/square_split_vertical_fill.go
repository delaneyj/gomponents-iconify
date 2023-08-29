package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareSplitVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 56v60a4 4 0 0 1-4 4H44a4 4 0 0 1-4-4V56a16 16 0 0 1 16-16h144a16 16 0 0 1 16 16Zm-4 80H44a4 4 0 0 0-4 4v60a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16v-60a4 4 0 0 0-4-4Z"/>`),
		g.Group(children),
	)
}