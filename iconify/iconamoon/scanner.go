package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4m12-8h2a2 2 0 0 1 2 2v2M8 20H6a2 2 0 0 1-2-2v-2m16 0v2a2 2 0 0 1-2 2h-2M4 8V6a2 2 0 0 1 2-2h2"/>`),
		g.Group(children),
	)
}