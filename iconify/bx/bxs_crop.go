package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsCrop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M19 7.5C19 6.121 17.879 5 16.5 5H8V2H5v3H2v3h14v14h3v-3h3v-3h-3V7.5z" fill="currentColor"/><path d="M8 10H5v6.5C5 17.879 6.121 19 7.5 19H14v-3H8v-6z" fill="currentColor"/>`),
		g.Group(children),
	)
}