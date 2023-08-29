package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuietTimeOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.1 23q-2.1 0-3.937-.8t-3.2-2.163Q4.6 18.675 3.8 16.837T3 12.9q0-3.2 1.8-5.8t4.825-3.65q.55-.2 1.025.137t.45.913q-.075 2.125.675 4.05t2.25 3.425q1.5 1.5 3.425 2.25t4.05.675q.65-.025.963.438t.112 1.037q-1.1 3-3.687 4.813T13.1 23Zm0-2q2.2 0 4.075-1.1t2.95-3.025q-2.15-.2-4.075-1.088t-3.45-2.412q-1.525-1.525-2.425-3.45T9.1 5.85Q7.175 6.925 6.088 8.813T5 12.9q0 3.375 2.363 5.738T13.1 21Zm-.5-7.625Z"/>`),
		g.Group(children),
	)
}