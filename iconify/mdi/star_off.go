package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StarOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m20.8 22.7l-2.9-2.9l.3 1.2l-6.2-3.7L5.8 21l1.6-7L2 9.2l4.9-.4L1.1 3l1.3-1.3l19.7 19.7l-1.3 1.3M22 9.2l-7.2-.6L12 2l-2 4.8l6.9 6.9L22 9.2Z"/>`),
		g.Group(children),
	)
}