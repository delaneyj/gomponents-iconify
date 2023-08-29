package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FridgeIndustrialOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 15.8L5.7 2.5c.35-.31.8-.5 1.3-.5h10a2 2 0 0 1 2 2v11.8m3.11 5.66l-1.27 1.27l-2.38-2.38c-.36.4-.88.65-1.46.65v1h-2v-1H9v1H7v-1a2 2 0 0 1-2-2V6.89L1.11 3l1.28-1.27l19.72 19.73M10 11.89L8.11 10H8v5h2v-3.11Z"/>`),
		g.Group(children),
	)
}