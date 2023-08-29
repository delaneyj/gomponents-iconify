package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.362 2.595l-1.671 3.05a9 9 0 0 1-4.785 4.122l-3.829 1.41l-.737 6.066L11 11.582l1.414 1.414l-5.66 5.66l6.063-.737l1.418-3.884a9 9 0 0 1 4.027-4.75l3.052-1.724l.984 1.741l-3.052 1.725a7 7 0 0 0-3.132 3.694l-1.838 5.036l-11.426 1.39L4.239 9.72l4.976-1.83a7 7 0 0 0 3.722-3.206l1.671-3.05l1.754.96Z"/>`),
		g.Group(children),
	)
}