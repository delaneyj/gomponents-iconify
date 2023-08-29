package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageInfoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.75 20.25q-1.575 0-2.663-1.088T14 16.5q0-1.575 1.088-2.663t2.662-1.087q1.575 0 2.663 1.088T21.5 16.5q0 1.575-1.088 2.663T17.75 20.25ZM11 17.5H5q-.425 0-.713-.288T4 16.5q0-.425.288-.713T5 15.5h6q.425 0 .713.288T12 16.5q0 .425-.288.713T11 17.5Zm-4.75-6.25q-1.575 0-2.663-1.088T2.5 7.5q0-1.575 1.088-2.663T6.25 3.75q1.575 0 2.663 1.088T10 7.5q0 1.575-1.088 2.663T6.25 11.25ZM19 8.5h-6q-.425 0-.713-.288T12 7.5q0-.425.288-.713T13 6.5h6q.425 0 .713.288T20 7.5q0 .425-.288.713T19 8.5Z"/>`),
		g.Group(children),
	)
}