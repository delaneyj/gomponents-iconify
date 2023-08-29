package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerGen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18q2.5 0 4.25-1.75T18 12q0-2.5-1.75-4.25T12 6Q9.5 6 7.75 7.75T6 12q0 2.5 1.75 4.25T12 18Zm0-2q-1.65 0-2.825-1.175T8 12q0-1.65 1.175-2.825T12 8q1.65 0 2.825 1.175T16 12q0 1.65-1.175 2.825T12 16Zm-.7-3.3q.275.275.7.275t.7-.275l1.4-1.4q.275-.275.275-.7t-.275-.7q-.275-.275-.7-.275t-.7.275l-1.4 1.4q-.275.275-.275.7t.275.7ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}