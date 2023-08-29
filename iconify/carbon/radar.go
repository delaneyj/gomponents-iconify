package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30 3.414L28.586 2L15.293 15.293a1 1 0 0 0 1.414 1.414l4.18-4.18A5.996 5.996 0 1 1 16 10V8a8.011 8.011 0 1 0 6.316 3.098l2.847-2.847A11.881 11.881 0 0 1 28 16A12 12 0 1 1 16 4V2a14 14 0 1 0 14 14a13.857 13.857 0 0 0-3.422-9.164Z"/>`),
		g.Group(children),
	)
}