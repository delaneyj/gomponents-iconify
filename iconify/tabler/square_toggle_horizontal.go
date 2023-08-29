package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareToggleHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12H2m2 2V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8m-2 6a2 2 0 0 0 2-2M4 18a2 2 0 0 0 2 2m8 0h-4"/>`),
		g.Group(children),
	)
}