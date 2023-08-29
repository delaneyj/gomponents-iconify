package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Visa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.756 6.342C4.344 5.56 2.733 4.93.931 4.494l.059-.35h7.407c.997.036 1.804.355 2.082 1.422l1.61 7.743v.001l.48 2.335l4.497-11.491h4.865L14.7 20.974l-4.86.005L5.756 6.342Z"/>`),
		g.Group(children),
	)
}