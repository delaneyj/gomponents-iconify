package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Highway(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 2L8 8h3V2h-1m3 0v6h3l-2-6h-1M2 9v1h2v1h2v-1h12l.06 1H20v-1h2V9H2m5 2L3.34 22H11V11H7m6 0v11h7.66L17 11h-4Z"/>`),
		g.Group(children),
	)
}