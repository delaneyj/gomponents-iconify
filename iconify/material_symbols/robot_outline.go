package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RobotOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V9q0-2.5 1.75-4.25T9 3h6q2.5 0 4.25 1.75T21 9v10q0 .825-.588 1.413T19 21H5Zm0-2h14V9q0-1.65-1.175-2.825T15 5H9Q7.35 5 6.175 6.175T5 9v10Zm4-7q-.825 0-1.413-.588T7 10q0-.825.588-1.413T9 8q.825 0 1.413.588T11 10q0 .825-.588 1.413T9 12Zm6 0q-.825 0-1.413-.588T13 10q0-.825.588-1.413T15 8q.825 0 1.413.588T17 10q0 .825-.588 1.413T15 12Zm-8 7v-2q0-.825.588-1.413T9 15h6q.825 0 1.413.588T17 17v2h-2v-2h-2v2h-2v-2H9v2H7Zm-2 0h14H5Z"/>`),
		g.Group(children),
	)
}