package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewGalleryOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 3v18h22V3H1m20 2v9H3V5h18M11 16v3H8v-3h3m-8 0h3v3H3v-3m10 3v-3h3v3h-3m5 0v-3h3v3h-3Z"/>`),
		g.Group(children),
	)
}