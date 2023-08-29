package heroicons_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.78 16.78a.75.75 0 0 1-1.06 0L4.5 5.56v7.69a.75.75 0 0 1-1.5 0v-9.5A.75.75 0 0 1 3.75 3h9.5a.75.75 0 0 1 0 1.5H5.56l11.22 11.22a.75.75 0 0 1 0 1.06Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}