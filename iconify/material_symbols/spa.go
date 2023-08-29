package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.35-.3-4.2-1.2t-3.138-2.35Q3.375 17 2.7 15.088t-.675-4.238q2.75.275 4.65 1t3.088 2.05q1.187 1.325 1.712 3.313T12 22Zm0-8.425q-.575-.875-1.563-1.725T8.15 10.3q.15-1.05.5-2.175t.85-2.212q.5-1.088 1.137-2.088T12 2q.725.825 1.363 1.825T14.5 5.912Q15 7 15.35 8.125t.5 2.175q-1.3.675-2.288 1.525T12 13.575Zm2 8.025q-.05-1.75-.263-3.238t-.662-2.812q1.175-2.025 3.238-3.3t5.662-1.4q.025 3.95-2.113 6.813T14 21.6Z"/>`),
		g.Group(children),
	)
}