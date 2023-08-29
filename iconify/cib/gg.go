package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m11.2 14.4l6.4 6.4l-6.4 6.4L0 16L11.2 4.8L14 7.6l-1.6 1.6L11.2 8l-8 8l8 8l3.219-3.219L9.6 16zm9.6-9.6l-6.4 6.4l6.4 6.4l1.6-1.6l-4.819-4.781L20.8 8l8 8l-8 8l-1.2-1.2l-1.6 1.6l2.8 2.8L32 16z"/>`),
		g.Group(children),
	)
}