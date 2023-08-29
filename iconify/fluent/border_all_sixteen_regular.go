package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAllSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.5A2.5 2.5 0 0 1 4.5 2h7A2.5 2.5 0 0 1 14 4.5v7a2.5 2.5 0 0 1-2.5 2.5h-7A2.5 2.5 0 0 1 2 11.5v-7ZM4.5 3A1.5 1.5 0 0 0 3 4.5v3h4.5V3h-3Zm4 0v4.5H13v-3A1.5 1.5 0 0 0 11.5 3h-3ZM13 8.5H8.5V13h3a1.5 1.5 0 0 0 1.5-1.5v-3ZM7.5 13V8.5H3v3A1.5 1.5 0 0 0 4.5 13h3Z"/>`),
		g.Group(children),
	)
}