package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReplyAllOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 9.8v.9l1.7.2c2.6.4 4.5 1.4 5.9 2.7c-1.7-.5-3.5-.8-5.6-.8h-2v1.3L8.8 12L11 9.8M13 5l-7 7l7 7v-4.1c5 0 8.5 1.6 11 5.1c-1-5-4-10-11-11M7 8V5l-7 7l7 7v-3l-4-4"/>`),
		g.Group(children),
	)
}