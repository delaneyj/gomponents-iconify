package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontalThirtyTwoFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 15a1 1 0 1 0 0 2h26a1 1 0 1 0 0-2H3Zm3 11.5v-8h20v8a3.5 3.5 0 0 1-3.5 3.5h-13A3.5 3.5 0 0 1 6 26.5Zm20-13v-8A3.5 3.5 0 0 0 22.5 2h-13A3.5 3.5 0 0 0 6 5.5v8h20Z"/>`),
		g.Group(children),
	)
}