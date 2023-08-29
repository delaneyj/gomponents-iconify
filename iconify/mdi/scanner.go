package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.2 10.7L19.8 5l.7 1.9L6.4 12H19a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-5.5c0-.8.5-1.6 1.2-1.8M17 17h2v-2h-2v2M5 17h10v-2H5v2Z"/>`),
		g.Group(children),
	)
}