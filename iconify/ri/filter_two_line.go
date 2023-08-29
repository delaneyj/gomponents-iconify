package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 14v6l-4 2v-8L4 5V3h16v2l-6 9ZM6.404 5L12 13.394L17.596 5H6.404Z"/>`),
		g.Group(children),
	)
}