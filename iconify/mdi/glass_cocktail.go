package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassCocktail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.5 7l-2-2h13l-2 2M11 13v6H6v2h12v-2h-5v-6l8-8V3H3v2l8 8Z"/>`),
		g.Group(children),
	)
}