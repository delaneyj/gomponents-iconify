package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UmbrellaOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4C8.9 4 6.18 6.03 5.3 9h13.4A7 7 0 0 0 12 4m0-2a9 9 0 0 1 9 9h-8v8a3 3 0 0 1-3 3a3 3 0 0 1-3-3v-1h2v1a1 1 0 0 0 1 1a1 1 0 0 0 1-1v-8H3a9 9 0 0 1 9-9Z"/>`),
		g.Group(children),
	)
}