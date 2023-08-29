package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpLeftLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 10v9h-2v-7H6.828l3.95 3.95l-1.414 1.414L3 11l6.364-6.364l1.414 1.414L6.828 10H19Z"/>`),
		g.Group(children),
	)
}