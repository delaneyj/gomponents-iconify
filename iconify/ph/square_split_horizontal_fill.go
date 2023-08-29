package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareSplitHorizontalFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 44v168a4 4 0 0 1-4 4H56a16 16 0 0 1-16-16V56a16 16 0 0 1 16-16h60a4 4 0 0 1 4 4Zm80-4h-60a4 4 0 0 0-4 4v168a4 4 0 0 0 4 4h60a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Z"/>`),
		g.Group(children),
	)
}