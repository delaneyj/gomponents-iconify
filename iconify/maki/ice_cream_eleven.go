package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IceCreamEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M4 6a2.49 2.49 0 0 0 1.5-.5c.432.325.959.5 1.5.5l-1 4.69a.5.5 0 0 1-.92 0L4 6zm3-4h-.09a1.5 1.5 0 1 0-2.82 0H4a1.5 1.5 0 1 0 1.5 1.5A1.5 1.5 0 1 0 7 2z" fill="currentColor"/>`),
		g.Group(children),
	)
}