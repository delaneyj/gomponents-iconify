package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-4h7.5c1.93 0 3.5-1.57 3.5-3.5V11h-4l6-7l6 7h-4v2.5c0 4.14-3.36 7.5-7.5 7.5H3Z"/>`),
		g.Group(children),
	)
}