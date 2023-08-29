package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TooltipCheckOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2H4c-1.1 0-2 .9-2 2v12a2 2 0 0 0 2 2h4l4 4l4-4h4c1.11 0 2-.89 2-2V4a2 2 0 0 0-2-2m0 14h-4.83L12 19.17L8.83 16H4V4h16v12m-9.25-2.29l-3.5-3.5l1.41-1.42l2.09 2.09l4.59-4.58l1.41 1.41l-6 6Z"/>`),
		g.Group(children),
	)
}