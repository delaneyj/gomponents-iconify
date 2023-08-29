package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BucketOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4h18v3h-1l-2.5 14h-11L4 7H3V4m14.97 3H6.03l2.12 12h7.7l2.12-12Z"/>`),
		g.Group(children),
	)
}