package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowTie(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15 14l6 3V7l-6 3v4m-6 0l-6 3V7l6 3v4m1-4h4v4h-4v-4Z"/>`),
		g.Group(children),
	)
}