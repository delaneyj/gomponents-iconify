package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedoVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20v-1h-1v1h1M10 8a6 6 0 0 0-6 6a6 6 0 0 0 6 6h4v-1h-4a5 5 0 0 1-5-5a5 5 0 0 1 5-5h7.09l-3.04 3.04l.71.7L19 8.5l-4.24-4.24l-.71.7L17.09 8H10Z"/>`),
		g.Group(children),
	)
}