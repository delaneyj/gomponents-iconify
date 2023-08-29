package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopologyStarTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 20a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm0-16a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm-8 8a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm16 0a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm-8 0a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm-8 0h4m4 0h4m-6-6v4m0 4v4"/>`),
		g.Group(children),
	)
}