package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FridgeIndustrialAlertOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 15H6v-5h2v5m9-11v15c0 1.11-.89 2-2 2v1h-2v-1H7v1H5v-1a2 2 0 0 1-2-2V4c0-1.1.9-2 2-2h10a2 2 0 0 1 2 2m-2 0H5v15h10V4m4 13h2v-2h-2v2m0-10v6h2V7h-2Z"/>`),
		g.Group(children),
	)
}