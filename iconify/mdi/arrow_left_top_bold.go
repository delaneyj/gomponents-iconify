package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLeftTopBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 21h-4v-7.5c0-1.93-1.57-3.5-3.5-3.5H11v4L4 8l7-6v4h2.5c4.14 0 7.5 3.36 7.5 7.5V21Z"/>`),
		g.Group(children),
	)
}