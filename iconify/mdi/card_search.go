package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 9a2.5 2.5 0 0 0 0 5a2.5 2.5 0 0 0 0-5M20 4H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2m-3.21 14.21l-2.91-2.91c-.69.44-1.51.7-2.38.7C9 16 7 14 7 11.5S9 7 11.5 7a4.481 4.481 0 0 1 3.8 6.89l2.91 2.9l-1.42 1.42Z"/>`),
		g.Group(children),
	)
}