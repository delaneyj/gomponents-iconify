package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondStone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 3h11l3.9 5.57L11.5 22L2.1 8.57L6 3m4.16 1l-1.7 4h6.08l-1.7-4h-2.68M8.33 9l3.17 9.76L14.67 9H8.33M3.72 8h3.66l1.7-4H6.5L3.72 8m-.1 1l6.82 9.75L7.28 9H3.62m15.66-1L16.5 4h-2.58l1.7 4h3.66m.1 1h-3.66l-3.16 9.75L19.38 9Z"/>`),
		g.Group(children),
	)
}