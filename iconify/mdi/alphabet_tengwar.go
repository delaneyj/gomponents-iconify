package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlphabetTengwar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 7l2.5-4H15l-4 4h-1m2.5 2c-.54 0-1.04.13-1.5.35V9H8v2h1v10h2v-8.5c0-.83.67-1.5 1.5-1.5s1.5.67 1.5 1.5v2c0 .83-.67 1.5-1.5 1.5H12v2h.5c1.93 0 3.5-1.57 3.5-3.5v-2c0-1.93-1.57-3.5-3.5-3.5Z"/>`),
		g.Group(children),
	)
}