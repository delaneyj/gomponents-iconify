package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7v11m0-4h18m0 4v-8a2 2 0 0 0-2-2h-8v6m-5-4a1 1 0 1 0 2 0a1 1 0 1 0-2 0"/>`),
		g.Group(children),
	)
}