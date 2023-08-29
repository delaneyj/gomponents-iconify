package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UltraHighDefinition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 7h2v4h2V7h2v10h-2v-4h-2v4H9V7m8 0h3a3 3 0 0 1 3 3v4a3 3 0 0 1-3 3h-3V7m3 8a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-1v6h1M7 14a3 3 0 0 1-3 3a3 3 0 0 1-3-3V7h2v7a1 1 0 0 0 1 1a1 1 0 0 0 1-1V7h2v7Z"/>`),
		g.Group(children),
	)
}