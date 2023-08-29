package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkMultipleThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.309 4H19.5A5.5 5.5 0 0 1 25 9.5v16.66a3.6 3.6 0 0 0 2-3.227V9.5A7.5 7.5 0 0 0 19.5 2h-6.967A3.6 3.6 0 0 0 9.31 4ZM8.75 6A3.75 3.75 0 0 0 5 9.75V29a1 1 0 0 0 1.591.806l6.945-5.092l7.92 5.126A1 1 0 0 0 23 29V9.75A3.75 3.75 0 0 0 19.25 6H8.75Z"/>`),
		g.Group(children),
	)
}