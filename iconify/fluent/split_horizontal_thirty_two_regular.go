package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 15a1 1 0 1 0 0 2h26a1 1 0 1 0 0-2H3Zm5 11.5v-8H6v8A3.5 3.5 0 0 0 9.5 30h13a3.5 3.5 0 0 0 3.5-3.5v-8h-2v8a1.5 1.5 0 0 1-1.5 1.5h-13A1.5 1.5 0 0 1 8 26.5Zm18-13v-8A3.5 3.5 0 0 0 22.5 2h-13A3.5 3.5 0 0 0 6 5.5v8h2v-8A1.5 1.5 0 0 1 9.5 4h13A1.5 1.5 0 0 1 24 5.5v8h2Z"/>`),
		g.Group(children),
	)
}