package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 3H2v13h19V3M2 17h4v4H2v-4m5 0h4v4H7v-4m5 0h4v4h-4v-4m5 0h4v4h-4v-4Z"/>`),
		g.Group(children),
	)
}