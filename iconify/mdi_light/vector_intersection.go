package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorIntersection(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h1v1H4a2 2 0 0 0-2 2v1H1V5a3 3 0 0 1 3-3m5 0v1H7V2h2m3 0a3 3 0 0 1 3 3v1h-1V5a2 2 0 0 0-2-2h-1V2h1m3 6v5a3 3 0 0 1-3 3H7v-5a3 3 0 0 1 3-3h5M2 10H1V8h1v2m0 3a2 2 0 0 0 2 2h1v1H4a3 3 0 0 1-3-3v-1h1v1m6 6a2 2 0 0 0 2 2h1v1h-1a3 3 0 0 1-3-3v-1h1v1m10 2a2 2 0 0 0 2-2v-1h1v1a3 3 0 0 1-3 3h-1v-1h1m2-10a2 2 0 0 0-2-2h-1V8h1a3 3 0 0 1 3 3v1h-1v-1m1 5h-1v-2h1v2m-8 6v-1h2v1h-2m-1-7a2 2 0 0 0 2-2V9h-4a2 2 0 0 0-2 2v4h4Z"/>`),
		g.Group(children),
	)
}