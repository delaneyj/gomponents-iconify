package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightKPlusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 15h2.5q.425 0 .713-.288T10.5 14v-4q0-.425-.288-.713T9.5 9H7q-.425 0-.713.288T6 10v4q0 .425.288.713T7 15Zm.5-1v-1.5H9V14H7.5Zm0-2.5V10H9v1.5H7.5ZM11 15h1.5v-2.25L14.25 15h1.825l-2.325-3l2.325-3H14.25l-1.75 2.25V9H11v6Zm5.5-1h1v-1.5H19v-1h-1.5V10h-1v1.5H15v1h1.5V14ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-2h14V5H5v14ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}