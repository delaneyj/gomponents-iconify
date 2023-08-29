package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PsychologyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 22v-4.3q-1.425-1.3-2.212-3.038T3 11q0-3.75 2.625-6.375T12 2q3.125 0 5.538 1.838t3.137 4.787l1.3 5.125q.125.475-.175.863T21 15h-2v3q0 .825-.588 1.413T17 20h-2v2h-2v-4h4v-5h2.7l-.95-3.875q-.575-2.275-2.45-3.7T12 4Q9.1 4 7.05 6.025T5 10.95q0 1.5.613 2.85t1.737 2.4l.65.6V22H6Zm6.35-9ZM11 15h2l.15-1.25q.2-.075.363-.175t.287-.225l1.15.5l1-1.7l-1-.75q.05-.2.05-.4t-.05-.4l1-.75l-1-1.7l-1.15.5q-.125-.125-.287-.225t-.363-.175L13 7h-2l-.15 1.25q-.2.075-.362.175t-.288.225l-1.15-.5l-1 1.7l1 .75Q9 10.8 9 11t.05.4l-1 .75l1 1.7l1.15-.5q.125.125.288.225t.362.175L11 15Zm1-2.5q-.625 0-1.062-.438T10.5 11q0-.625.438-1.063T12 9.5q.625 0 1.063.438T13.5 11q0 .625-.438 1.063T12 12.5Z"/>`),
		g.Group(children),
	)
}