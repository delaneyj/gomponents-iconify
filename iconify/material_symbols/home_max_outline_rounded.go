package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeMaxOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17h14q.825 0 1.413-.588T21 15V9q0-.825-.588-1.413T19 7H5q-.825 0-1.413.588T3 9v6q0 .825.588 1.413T5 17Zm0 2q-1.65 0-2.825-1.175T1 15V9q0-1.65 1.175-2.825T5 5h14q1.65 0 2.825 1.175T23 9v6q0 1.65-1.175 2.825T19 19h-2q0 .425-.288.713T16 20H8q-.425 0-.713-.288T7 19H5Zm7-7Z"/>`),
		g.Group(children),
	)
}