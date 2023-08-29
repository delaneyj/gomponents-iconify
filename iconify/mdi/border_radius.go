package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderRadius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16c0 2.8 2.2 5 5 5h2v-2H8c-1.7 0-3-1.3-3-3v-2H3v2m18-8c0-2.8-2.2-5-5-5h-2v2h2c1.7 0 3 1.3 3 3v2h2V8m-5 13c2.8 0 5-2.2 5-5v-2h-2v2c0 1.7-1.3 3-3 3h-2v2h2M8 3C5.2 3 3 5.2 3 8v2h2V8c0-1.7 1.3-3 3-3h2V3H8Z"/>`),
		g.Group(children),
	)
}