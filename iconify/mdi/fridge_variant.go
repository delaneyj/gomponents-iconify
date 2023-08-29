package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FridgeVariant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 4v15c0 1.11-.89 2-2 2v1h-2v-1h-2.5V2H17a2 2 0 0 1 2 2M7 2h4.5v19H9v1H7v-1a2 2 0 0 1-2-2V4c0-1.1.9-2 2-2m3 8H7v4h3v-4Z"/>`),
		g.Group(children),
	)
}