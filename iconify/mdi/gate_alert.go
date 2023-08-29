package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GateAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 17h-2v2h2v-2m0-8h-2v6h2V9m-4 4v-2h-2V6h-2v5h-2V6H9v5H7V7H5v4H3V9H1v12h2v-2h2v2h2v-2h2v2h2v-2h2v2h2v-2h2v-2h-2v-4h2M5 17H3v-4h2v4m4 0H7v-4h2v4m4 0h-2v-4h2v4Z"/>`),
		g.Group(children),
	)
}