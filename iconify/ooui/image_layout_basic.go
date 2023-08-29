package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageLayoutBasic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 3v14h18V3zm17 13H2V4h16z"/><path fill="currentColor" d="M8.58 14h.81l3.11-4l3 4H17l-4.5-6L9 12.51L6.5 9.5L3 14h1.56l1.94-2.5z"/>`),
		g.Group(children),
	)
}