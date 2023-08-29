package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrpanoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 15.4q1.625-.2 3.238-.3T12 15q1.65 0 3.263.1t3.237.3L14 10l-2.85 3.4l-2-2.4l-3.65 4.4ZM3 20q-.425 0-.713-.288T2 19V5q0-.425.288-.713T3 4q.2 0 .888.238t1.837.512q1.15.275 2.713.513T12 5.5q2 0 3.563-.238t2.712-.512q1.15-.275 1.838-.513T21 4q.425 0 .713.288T22 5v14q0 .425-.288.713T21 20q-.2 0-.888-.238t-1.837-.512q-1.15-.275-2.712-.513T12 18.5q-2 0-3.563.238t-2.712.512q-1.15.275-1.837.513T3 20Zm1-2.35q1.95-.575 3.963-.863T12 16.5q2.025 0 4.038.288T20 17.65V6.375q-1.95.575-3.963.85T12 7.5q-2.025 0-4.038-.275T4 6.375V17.65ZM12 12Z"/>`),
		g.Group(children),
	)
}