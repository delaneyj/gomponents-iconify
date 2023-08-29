package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalHspaPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 8v3h3v3h-3v3h-3v-3h-3v-3h3V8h3M5 10.5h3v-6h3v15H8v-6H5v6H2v-15h3v6Z"/>`),
		g.Group(children),
	)
}