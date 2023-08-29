package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRotateCounterclockwiseTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 3.5a6.5 6.5 0 0 1 6.5 6.5a.75.75 0 0 0 1.5 0a8 8 0 1 0-12.664 6.5H4.25a.75.75 0 0 0 0 1.5h3a.75.75 0 0 0 .75-.75v-3a.75.75 0 0 0-1.5 0v1.228A6.5 6.5 0 0 1 10 3.5ZM7.5 10a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0ZM9 10a1 1 0 1 1 2 0a1 1 0 0 1-2 0Z"/>`),
		g.Group(children),
	)
}