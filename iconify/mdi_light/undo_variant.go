package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UndoVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 20v-1h1v1H6m7-12a6 6 0 0 1 6 6a6 6 0 0 1-6 6H9v-1h4a5 5 0 0 0 5-5a5 5 0 0 0-5-5H5.91l3.04 3.04l-.71.7L4 8.5l4.24-4.24l.71.7L5.91 8H13Z"/>`),
		g.Group(children),
	)
}