package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PanoramaHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20q-.425 0-.713-.288T2 19V5q0-.425.288-.713T3 4q.2 0 .888.238t1.837.512q1.15.275 2.713.513T12 5.5q2 0 3.563-.238t2.712-.512q1.15-.275 1.838-.513T21 4q.425 0 .713.288T22 5v14q0 .425-.288.713T21 20q-.2 0-.888-.238t-1.837-.512q-1.15-.275-2.712-.513T12 18.5q-2 0-3.563.238t-2.712.512q-1.15.275-1.837.513T3 20Z"/>`),
		g.Group(children),
	)
}