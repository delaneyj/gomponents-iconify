package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VectorDifferenceBa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 2h8a3 3 0 0 1 3 3v4h-5a2 2 0 0 0-2 2v5H4a3 3 0 0 1-3-3V5a3 3 0 0 1 3-3m10 6V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h3v-4a3 3 0 0 1 3-3h4m4 0a3 3 0 0 1 3 3v1h-1v-1a2 2 0 0 0-2-2h-1V8h1m3 8h-1v-2h1v2m0 3a3 3 0 0 1-3 3h-1v-1h1a2 2 0 0 0 2-2v-1h1v1m-8 3v-1h2v1h-2m-3 0a3 3 0 0 1-3-3v-1h1v1a2 2 0 0 0 2 2h1v1h-1Z"/>`),
		g.Group(children),
	)
}