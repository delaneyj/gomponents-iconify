package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardFNine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 7h6v2H7v2h3v2H7v4H5V7m12 10h-4v-2h4v-2h-2a2 2 0 0 1-2-2V9c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v6c0 1.11-.89 2-2 2m0-6V9h-2v2h2Z"/>`),
		g.Group(children),
	)
}