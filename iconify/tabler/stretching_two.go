package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StretchingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4a1 1 0 1 0 2 0a1 1 0 0 0-2 0M6.5 21l3.5-5m-5-5l7-2m4 12l-4-7V9l7-4"/>`),
		g.Group(children),
	)
}