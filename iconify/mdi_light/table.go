package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 5h11a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V8a3 3 0 0 1 3-3M4 17a2 2 0 0 0 2 2h5v-3H4v1m7-5H4v3h7v-3m6 7a2 2 0 0 0 2-2v-1h-7v3h5m2-7h-7v3h7v-3M4 11h7V8H4v3m8 0h7V8h-7v3Z"/>`),
		g.Group(children),
	)
}