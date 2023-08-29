package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SourceNotesOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19V5v14ZM7 9h10V7H7v2Zm0 4h4.75q.5-.6 1.088-1.113T14.125 11H7v2Zm0 4h3.05q.05-.525.188-1.025t.337-.975H7v2Zm-4 4V3h18v7.575q-.475-.2-.975-.338T19 10.05V5H5v14h5.05q.05.525.188 1.025t.337.975H3Zm15 2q-1.825 0-3.188-1.137T13.1 19h1.55q.325 1.1 1.238 1.8t2.112.7q1.45 0 2.475-1.025T21.5 18q0-1.45-1.025-2.475T18 14.5q-.725 0-1.35.263t-1.1.737H17V17h-4v-4h1.5v1.425q.675-.65 1.575-1.038T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Z"/>`),
		g.Group(children),
	)
}