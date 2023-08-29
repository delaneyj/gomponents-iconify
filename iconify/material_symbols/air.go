package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Air(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 20q-1.25 0-2.125-.875T8.5 17h2q0 .425.288.713T11.5 18q.425 0 .713-.288T12.5 17q0-.425-.288-.713T11.5 16H2v-2h9.5q1.25 0 2.125.875T14.5 17q0 1.25-.875 2.125T11.5 20ZM2 10V8h13.5q.65 0 1.075-.425T17 6.5q0-.65-.425-1.075T15.5 5q-.65 0-1.075.425T14 6.5h-2q0-1.475 1.012-2.488T15.5 3q1.475 0 2.488 1.012T19 6.5q0 1.475-1.012 2.488T15.5 10H2Zm16.5 8v-2q.65 0 1.075-.425T20 14.5q0-.65-.425-1.075T18.5 13H2v-2h16.5q1.475 0 2.488 1.012T22 14.5q0 1.475-1.012 2.488T18.5 18Z"/>`),
		g.Group(children),
	)
}