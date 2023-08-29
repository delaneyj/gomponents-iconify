package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.272 6.008A2 2 0 0 0 20 5a2 2 0 1 0-3.728 1.008Zm0 0l-8.544 4.984m0 0A1.999 1.999 0 0 0 4 12a2 2 0 0 0 3.728 1.008m0-2.016a2 2 0 0 1 0 2.016m0 0l8.544 4.984m0 0A1.999 1.999 0 0 1 20 19a2 2 0 1 1-3.728-1.008Z"/>`),
		g.Group(children),
	)
}