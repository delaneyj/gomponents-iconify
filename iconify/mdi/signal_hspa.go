package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalHspa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 10.5h3v-6h3v15h-3v-6h-3v6h-3v-15h3v6Z"/>`),
		g.Group(children),
	)
}