package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NightsStay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 14q1.2 0 2.2.65t1.475 1.775l.25.575h.625q1.05 0 1.75.738T13 19.5q0 1.05-.725 1.775T10.5 22H6q-1.65 0-2.825-1.175T2 18q0-1.675 1.175-2.838T6 14Zm5.25-11q-.45 2.475.275 4.838t2.5 4.137q1.775 1.775 4.138 2.5T23 14.75q-.65 3.55-3.375 5.863T13.325 23q.8-.65 1.238-1.563T15 19.5q0-1.7-1.063-2.938t-2.712-1.487q-.8-1.425-2.187-2.25T6 12q-.8 0-1.563.2T3 12.8q.05-3.625 2.363-6.375T11.25 3Z"/>`),
		g.Group(children),
	)
}