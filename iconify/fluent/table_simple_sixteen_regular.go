package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 4.5v7A2.5 2.5 0 0 0 4.5 14h7a2.5 2.5 0 0 0 2.5-2.5v-7A2.5 2.5 0 0 0 11.5 2h-7A2.5 2.5 0 0 0 2 4.5Zm1 4h4.5V13h-3A1.5 1.5 0 0 1 3 11.5v-3Zm4.5-1H3v-3A1.5 1.5 0 0 1 4.5 3h3v4.5Zm1 1H13v3a1.5 1.5 0 0 1-1.5 1.5h-3V8.5Zm4.5-1H8.5V3h3A1.5 1.5 0 0 1 13 4.5v3Z"/>`),
		g.Group(children),
	)
}