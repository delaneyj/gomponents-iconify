package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LayersAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2H9a1 1 0 0 0-1 1v4H6a1 1 0 0 0-1 1v4H3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-2h4a1 1 0 0 0 1-1v-2h4a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1ZM10 20H4v-6h6Zm5-3h-3v-4a1 1 0 0 0-1-1H7V9h8Zm5-3h-3V8a1 1 0 0 0-1-1h-6V4h10Z"/>`),
		g.Group(children),
	)
}