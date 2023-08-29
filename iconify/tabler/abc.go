package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Abc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 16v-6a2 2 0 1 1 4 0v6m-4-3h4m3-5v6a2 2 0 1 0 4 0v-1a2 2 0 1 0-4 0v1m10.732-2A2 2 0 0 0 17 13v1a2 2 0 0 0 3.726 1.01"/>`),
		g.Group(children),
	)
}