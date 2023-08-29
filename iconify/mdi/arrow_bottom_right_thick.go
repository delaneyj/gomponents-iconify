package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBottomRightThick(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14.89 12.06V7.11h3.42v11.2H7.11v-3.42h4.95L5.69 8.5L8.5 5.69l6.39 6.37Z"/>`),
		g.Group(children),
	)
}