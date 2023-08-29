package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopologyRingTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 6a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM7 18a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm14 0a2 2 0 1 0-4 0a2 2 0 0 0 4 0zM7 18h10m1-2l-5-8m-2 0l-5 8"/>`),
		g.Group(children),
	)
}