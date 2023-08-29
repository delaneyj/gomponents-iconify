package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardTabReverse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6H2v12h2m7-12l-6 6l6 6l1.41-1.42L8.83 13H23v-2H8.83l3.58-3.59L11 6Z"/>`),
		g.Group(children),
	)
}