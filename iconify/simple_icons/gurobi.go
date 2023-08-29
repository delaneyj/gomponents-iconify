package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gurobi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.036 0l7.032 1.359L24 18.37L18.37 24L0 17.635L1.805 5.952L11.036 0Zm12.389 18.239L17.887 2.36l-3.557 7.83l3.88 13.264l5.215-5.214Zm-5.822-16.46L11.138.528l-8.71 5.617l11.554 3.6l3.62-7.968Z"/>`),
		g.Group(children),
	)
}