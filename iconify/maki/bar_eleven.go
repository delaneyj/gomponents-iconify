package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M5.488 1C4.976 1 .5 1 1 1.5L5 6v2.5C5 9 2.5 9 2.5 10h6C8.5 9 6 9 6 8.5V6l4-4.5C10.5 1 6 1 5.488 1zM2.5 2h6l-1 1h-4l-1-1z" fill="currentColor"/>`),
		g.Group(children),
	)
}