package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerDownLeftLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14V5h-2v7H6.828l3.95-3.95l-1.414-1.414L3 13l6.364 6.364l1.414-1.414L6.828 14H19Z"/>`),
		g.Group(children),
	)
}