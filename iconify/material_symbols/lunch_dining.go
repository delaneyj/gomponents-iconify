package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LunchDining(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 10V9q0-2.875 2.713-4.438T12 3q4.575 0 7.288 1.563T22 9v1H2Zm0 4.5v-2q.9 0 1.425-.5t1.925-.5q1.4 0 1.9.5t1.4.5q.9 0 1.425-.5T12 11.5q1.4 0 1.925.5t1.425.5q.9 0 1.4-.5t1.9-.5q1.4 0 1.975.5t1.375.5v2q-1.4 0-1.875-.5t-1.375-.5q-.9 0-1.45.5t-1.95.5q-1.4 0-1.925-.5T12 13.5q-.9 0-1.425.5t-1.925.5q-1.4 0-1.9-.5t-1.4-.5q-.9 0-1.425.5T2 14.5ZM4 21q-.825 0-1.413-.588T2 19v-3h20v3q0 .825-.588 1.413T20 21H4Z"/>`),
		g.Group(children),
	)
}