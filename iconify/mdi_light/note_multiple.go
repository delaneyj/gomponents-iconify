package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteMultiple(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 12a3 3 0 0 1-3-3V5H6a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h13a2 2 0 0 0 2-2v-6h-4m-2-3a2 2 0 0 0 2 2h3.59L15 5.41V9M6 4h9l7 7v7a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3m0 19a5 5 0 0 1-5-5V8h1v10a4 4 0 0 0 4 4h12v1H6Z"/>`),
		g.Group(children),
	)
}