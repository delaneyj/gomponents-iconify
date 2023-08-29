package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FridgeVariantAlert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 4v15c0 1.11-.89 2-2 2v1h-2v-1h-2.5V2H15a2 2 0 0 1 2 2M5 2h4.5v19H7v1H5v-1a2 2 0 0 1-2-2V4c0-1.1.9-2 2-2m3 8H5v4h3v-4m11-3v6h2V7h-2m0 10h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}