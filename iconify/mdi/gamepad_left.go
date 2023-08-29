package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamepadLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 2v5.5l3 3l3-3V2H9M2 9v6h5.5l3-3l-3-3H2m14.5 0l-3 3l3 3H22V9h-5.5M4 11h2v2H4v-2m8 2.5l-3 3V22h6v-5.5l-3-3Z"/>`),
		g.Group(children),
	)
}