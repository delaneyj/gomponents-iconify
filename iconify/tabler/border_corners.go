package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderCorners(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 4h2a2 2 0 0 1 2 2v2m0 8v2a2 2 0 0 1-2 2h-2m-8 0H6a2 2 0 0 1-2-2v-2m0-8V6a2 2 0 0 1 2-2h2"/>`),
		g.Group(children),
	)
}