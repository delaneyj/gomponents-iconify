package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipMinusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v12a2 2 0 0 0 2 2h4l4 4l4-4h4c1.11 0 2-.89 2-2V4a2 2 0 0 0-2-2m0 14h-4.83L12 19.17L8.83 16H4V4h16v12m-4-5H8V9h8v2Z"/>`),
		g.Group(children),
	)
}