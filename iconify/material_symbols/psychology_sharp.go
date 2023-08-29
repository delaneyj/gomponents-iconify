package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PsychologySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 15h2l.15-1.25q.2-.075.363-.175t.287-.225l1.15.5l1-1.7l-1-.75q.05-.2.05-.4t-.05-.4l1-.75l-1-1.7l-1.15.5q-.125-.125-.288-.225t-.362-.175L13 7h-2l-.15 1.25q-.2.075-.363.175t-.287.225l-1.15-.5l-1 1.7l1 .75Q9 10.8 9 11t.05.4l-1 .75l1 1.7l1.15-.5q.125.125.288.225t.362.175L11 15Zm1-2.5q-.625 0-1.063-.438T10.5 11q0-.625.438-1.063T12 9.5q.625 0 1.063.438T13.5 11q0 .625-.438 1.063T12 12.5ZM6 22v-4.3q-1.425-1.3-2.212-3.038T3 11q0-3.75 2.625-6.375T12 2q3.125 0 5.538 1.838t3.137 4.787l1.6 6.375H19v5h-4v2H6Z"/>`),
		g.Group(children),
	)
}