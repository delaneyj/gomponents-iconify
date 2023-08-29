package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HierarchyThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 5a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-4 7a2 2 0 1 0 4 0a2 2 0 1 0-4 0m4 7a2 2 0 1 0 4 0a2 2 0 1 0-4 0m8 0a2 2 0 1 0 4 0a2 2 0 1 0-4 0M2 19a2 2 0 1 0 4 0a2 2 0 1 0-4 0m12-7a2 2 0 1 0 4 0a2 2 0 1 0-4 0m-9 5l2-3m2-4l2-3m2 0l2 3m2 4l2 3m-4-3l-2 3m-4-3l2 3"/>`),
		g.Group(children),
	)
}