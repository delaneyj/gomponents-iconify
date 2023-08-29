package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TempleBuddhistSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.25 6L12 1l3.75 5h-7.5ZM6 10V8.85q-1.3-.325-2.15-1.375T3 5.025h2q0 .8.588 1.387T6.975 7h10.05q.8 0 1.388-.588T19 5.025h2q0 1.4-.85 2.45T18 8.85V10H6ZM4 22v-9.15q-1.3-.325-2.15-1.375T1 9.025h2q0 .8.588 1.388T4.974 11h14.05q.8 0 1.388-.588T21 9.025h2q0 1.4-.85 2.45T20 12.85V22h-7v-5h-2v5H4Z"/>`),
		g.Group(children),
	)
}