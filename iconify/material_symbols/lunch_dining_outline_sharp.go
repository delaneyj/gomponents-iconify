package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LunchDiningOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 10V9q0-2.875 2.713-4.438T12 3q4.575 0 7.288 1.563T22 9v1H2Zm2.15-2h15.7q-.575-1.45-2.663-2.225T12 5q-3.1 0-5.188.775T4.15 8ZM2 14.5v-2q.9 0 1.425-.5t1.925-.5q1.4 0 1.9.5t1.4.5q.9 0 1.425-.5T12 11.5q1.4 0 1.925.5t1.425.5q.9 0 1.4-.5t1.9-.5q1.4 0 1.975.5t1.375.5v2q-1.4 0-1.875-.5t-1.375-.5q-.9 0-1.45.5t-1.95.5q-1.4 0-1.925-.5T12 13.5q-.9 0-1.425.5t-1.925.5q-1.4 0-1.9-.5t-1.4-.5q-.9 0-1.425.5T2 14.5ZM2 21v-5h20v5H2Zm2-2h16v-1H4v1Zm.15-11h15.7h-15.7ZM4 18h16H4Z"/>`),
		g.Group(children),
	)
}