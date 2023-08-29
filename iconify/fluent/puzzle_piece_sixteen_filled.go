package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PuzzlePieceSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9 1a2 2 0 0 0-2 2H5.5A1.5 1.5 0 0 0 4 4.5V6a2 2 0 0 0 0 4v1.5A1.5 1.5 0 0 0 5.5 13H7a2 2 0 0 0 4 0h1.5a1.5 1.5 0 0 0 1.5-1.5V9h-1a1 1 0 1 1 0-2h1V4.5A1.5 1.5 0 0 0 12.5 3H11a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}