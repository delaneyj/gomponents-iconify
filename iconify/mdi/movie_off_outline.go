package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovieOffOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.39 1.73L1.11 3l1.54 1.54C2.25 4.9 2 5.42 2 6v12a2 2 0 0 0 2 2h14.11l2.73 2.73l1.27-1.27L2.39 1.73M4 18V6.47L5.76 10h2.35l8 8H4M8.8 5.6L8 4h2l2 4h3l-2-4h2l2 4h3l-2-4h4v14c0 .24-.04.47-.12.68L20 16.8V10h-6.8L8.8 5.6Z"/>`),
		g.Group(children),
	)
}