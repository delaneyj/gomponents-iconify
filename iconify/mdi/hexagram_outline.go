package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagramOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 7l1.3 2.8l3.2-.4l-2 2.6l1.9 2.5l-3.2-.4L12 17l-1.3-2.8l-3.2.4l2-2.6l-1.9-2.5l3.2.4L12 7m0-5L9.5 7.7L3.3 7L7 12l-3.6 5l6.2-.7L12 22l2.5-5.7l6.2.6L17 12l3.6-5l-6.2.7L12 2Z"/>`),
		g.Group(children),
	)
}