package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpLeftDoubleLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 10v9h-2v-7h-6.172l3.95 3.95l-1.414 1.414L8 11l6.364-6.364l1.414 1.414l-3.95 3.95H20ZM8.75 4.636l1.414 1.414L5.214 11l4.95 4.95l-1.414 1.414L2.386 11L8.75 4.636Z"/>`),
		g.Group(children),
	)
}