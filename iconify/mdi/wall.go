package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16h9v5H3v-5m-1-6h6v5H2v-5m7 0h6v5H9v-5m7 0h6v5h-6v-5m-3 6h8v5h-8v-5M3 4h8v5H3V4m9 0h9v5h-9V4Z"/>`),
		g.Group(children),
	)
}