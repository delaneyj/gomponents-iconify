package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Exponent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m15.38 3l2.39 5.75c-.22.93-.5 1.57-.77 1.95c-.33.48-.56.55-.81.55v1.5c.75 0 1.55-.4 2.05-1.19C19.87 8.94 22 3 22 3h-1.62l-1.69 4.05L17 3h-1.62M3.42 8.59L2 10l4.79 4.79L2 19.59L3.41 21l4.8-4.79L13 21l1.41-1.41l-4.79-4.8L14.41 10L13 8.59l-4.79 4.79l-4.8-4.79h.01Z"/>`),
		g.Group(children),
	)
}