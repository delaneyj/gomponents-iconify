package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrinterPosNetwork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 7H7V2h10v5m5 13v2h-7c0 .55-.45 1-1 1h-4c-.55 0-1-.45-1-1H2v-2h7c0-.55.45-1 1-1h1v-2H4v-7a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v7h-7v2h1c.55 0 1 .45 1 1h7m-4-10h-4v2h4v-2Z"/>`),
		g.Group(children),
	)
}