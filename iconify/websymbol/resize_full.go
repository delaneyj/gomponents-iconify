package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResizeFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 1v467L832 307L671 468q-29 29-69 29t-69-29t-29-69t29-69l161-161L533 1h467zM306 833l161 168H0V534l168 161l161-161q29-29 69-29t69 29t29 69t-29 69z"/>`),
		g.Group(children),
	)
}