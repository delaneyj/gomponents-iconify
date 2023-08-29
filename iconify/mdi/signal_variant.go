package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 6V4h.1C12.9 4 20 11.1 20 19.9v.1h-2v-.1C18 12.2 11.8 6 4 6m0 4V8a12 12 0 0 1 12 12h-2A10 10 0 0 0 4 10m0 4v-2a8 8 0 0 1 8 8h-2a6 6 0 0 0-6-6m0 2a4 4 0 0 1 4 4H4v-4Z"/>`),
		g.Group(children),
	)
}