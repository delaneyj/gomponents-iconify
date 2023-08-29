package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableSimpleExcludeThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.5 2A4.5 4.5 0 0 0 2 6.5V19a4.5 4.5 0 0 0 4.5 4.5H13a1 1 0 0 0 1-1V14h8.5a1 1 0 0 0 1-1V6.5A4.5 4.5 0 0 0 19 2H6.5ZM12 14v7.5H6.5A2.5 2.5 0 0 1 4 19v-5h8Zm0-2H4V6.5A2.5 2.5 0 0 1 6.5 4H12v8Zm2 0V4h5a2.5 2.5 0 0 1 2.5 2.5V12H14Zm2 7a3 3 0 0 1 3-3h8a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3h-8a3 3 0 0 1-3-3v-8Z"/>`),
		g.Group(children),
	)
}