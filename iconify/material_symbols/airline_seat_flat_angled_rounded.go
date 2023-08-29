package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatFlatAngledRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m9.4 11.775l1.375-3.75Q11.05 7.25 11.8 6.9t1.525-.075l6.575 2.4q1.575.575 2.275 2.063t.125 3.062l-.675 1.875q-.15.4-.525.575t-.775.025L10 13.075q-.4-.15-.575-.525t-.025-.775Zm9.875 7.875L2.35 13.5q-.4-.125-.562-.5t-.038-.775q.125-.4.513-.575t.787-.025l16.925 6.15q.4.125.563.5t.037.775q-.125.4-.512.575t-.788.025Zm-12.95-8.1q-1.25 0-2.125-.875T3.325 8.55q0-1.25.875-2.125t2.125-.875q1.25 0 2.125.875t.875 2.125q0 1.25-.875 2.125t-2.125.875Z"/>`),
		g.Group(children),
	)
}