package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wardrobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 4v15c0 1.1.9 2 2 2v1h2v-1h3.5V2H6c-1.1 0-2 .9-2 2m4 6h2v3H8v-3m10-8h-5.5v19H16v1h2v-1c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2m-2 11h-2v-3h2v3Z"/>`),
		g.Group(children),
	)
}