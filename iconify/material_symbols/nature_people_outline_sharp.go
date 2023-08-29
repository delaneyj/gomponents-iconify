package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NaturePeopleOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.5 11q-.65 0-1.075-.425T3 9.5q0-.65.425-1.075T4.5 8q.65 0 1.075.425T6 9.5q0 .65-.425 1.075T4.5 11ZM3 22v-5H2v-5h5v5H6v3h8v-5h-1.75q-1.775 0-3.013-1.238T8 10.75q0-1.325.713-2.363T10.55 6.85q.275-1.625 1.513-2.737T15 3q1.7 0 2.938 1.113T19.45 6.85q1.125.5 1.838 1.538T22 10.75q0 1.775-1.238 3.013T17.75 15H16v5h5v2H3Zm9.25-9h5.5q.95 0 1.6-.65t.65-1.6q0-.675-.363-1.225T18.65 8.7l-1.05-.45l-.15-1.1q-.15-.925-.837-1.537T15 5q-.925 0-1.613.613T12.55 7.15l-.15 1.1l-1.05.45q-.625.275-.988.825T10 10.75q0 .95.65 1.6t1.6.65ZM15 9Z"/>`),
		g.Group(children),
	)
}