package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxOneLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 1l9.5 5.5v11L12 23l-9.5-5.5v-11L12 1ZM5.494 7.078L13 11.423v8.687l6.5-3.763V7.653L12 3.311L5.494 7.078ZM4.5 8.813v7.534L11 20.11v-7.534L4.5 8.813Z"/>`),
		g.Group(children),
	)
}