package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 9.8v.9l1.7.3c2.6.4 4.5 1.4 5.9 2.7c-1.7-.5-3.5-.8-5.6-.8H8v1.3L5.8 12L8 9.8M10 5l-7 7l7 7v-4.1c5 0 8.5 1.6 11 5.1c-1-5-4-10-11-11"/>`),
		g.Group(children),
	)
}