package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7h4a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H5v-2h4v-2H6v-2h3V9H5V7m8 0h3a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3h-3V7m3 8a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-1v6h1Z"/>`),
		g.Group(children),
	)
}