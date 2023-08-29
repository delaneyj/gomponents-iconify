package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareRootTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 12h1c1 0 1 1 2.016 3.527C17 18 17 19 18 19h1"/><path d="M12 19c1.5 0 3-2 4-3.5s2.5-3.5 4-3.5M3 12h1l3 8l3-16h10"/></g>`),
		g.Group(children),
	)
}