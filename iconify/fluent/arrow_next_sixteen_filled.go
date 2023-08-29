package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowNextSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3.704 11.736a.75.75 0 1 0 1.092 1.028l4-4.25a.748.748 0 0 0 0-1.027l-4-4.25a.75.75 0 1 0-1.092 1.028L7.22 8l-3.516 3.736ZM11.25 3a.75.75 0 0 1 .75.75v8.5a.75.75 0 0 1-1.5 0v-8.5a.75.75 0 0 1 .75-.75Z"/>`),
		g.Group(children),
	)
}