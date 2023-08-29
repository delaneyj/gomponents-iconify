package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Psychiatry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-7.975q-1.6 0-3.075-.613t-2.6-1.737Q4.2 9.55 3.6 8.075T3 5V3h2q1.575 0 3.05.613t2.6 1.737q.775.775 1.288 1.7t.787 1.975q.125-.175.275-.338t.325-.337q1.125-1.125 2.6-1.738T19 6h2v2q0 1.6-.613 3.075t-1.737 2.6Q17.525 14.8 16.063 15.4T13 16v5h-2Z"/>`),
		g.Group(children),
	)
}