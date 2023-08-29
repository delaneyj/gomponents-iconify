package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attachment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 17.5q-2.3 0-3.9-1.6T2 12q0-2.3 1.6-3.9t3.9-1.6H18q1.65 0 2.825 1.175T22 10.5q0 1.65-1.175 2.825T18 14.5H8.5q-1.05 0-1.775-.725T6 12q0-1.05.725-1.775T8.5 9.5H18V11H8.5q-.425 0-.713.288T7.5 12q0 .425.288.713T8.5 13H18q1.05 0 1.775-.725T20.5 10.5q0-1.05-.725-1.775T18 8H7.5Q5.85 8 4.675 9.175T3.5 12q0 1.65 1.175 2.825T7.5 16H18v1.5H7.5Z"/>`),
		g.Group(children),
	)
}