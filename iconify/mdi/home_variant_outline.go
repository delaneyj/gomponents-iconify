package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeVariantOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 13h6v6h3v-9l-6-4.5L6 10v9h3v-6m-5 8V9l8-6l8 6v12H4Z"/>`),
		g.Group(children),
	)
}