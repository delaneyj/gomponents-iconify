package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxTwoLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l9.5 5.5v11L12 23l-9.5-5.5v-11L12 1ZM4.5 7.653v.005L12 12v8.689l7.5-4.342V7.653L12 3.311L4.5 7.653Z"/>`),
		g.Group(children),
	)
}