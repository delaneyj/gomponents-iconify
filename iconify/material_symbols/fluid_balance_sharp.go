package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluidBalanceSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v2H11v3H6v2h5v2H6v2h5q0 1.425.563 2.725T13.15 20H2Zm15 3v-3.1q-1.75-.35-2.875-1.725T13 15V8h10v7q0 1.8-1.125 3.175T19 19.9V21h3v2h-5Zm2.75-9H21v-4h-6v2h.75q.825 0 1.563.375T18.55 13.4q.2.3.525.45t.675.15Z"/>`),
		g.Group(children),
	)
}