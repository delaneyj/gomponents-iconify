package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.126 2.219a3 3 0 0 1 3.748 0l7.63 6.104a3 3 0 0 1 .945 3.367l-3.03 8.335A3 3 0 0 1 16.599 22H7.401a3 3 0 0 1-2.82-1.975l-3.03-8.334a3 3 0 0 1 .945-3.368l7.63-6.104Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}