package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 21V7H5v4H3V9H1v12h2v-2h2v2h2m-4-4v-4h2v4H3m18-8v2h-2V7h-2v14h2v-2h2v2h2V9h-2m0 8h-2v-4h2v4Z"/>`),
		g.Group(children),
	)
}