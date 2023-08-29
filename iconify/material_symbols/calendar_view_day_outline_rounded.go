package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewDayOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17q-.825 0-1.413-.588T3 15V9q0-.825.588-1.413T5 7h14q.825 0 1.413.588T21 9v6q0 .825-.588 1.413T19 17H5Zm0-2h14V9H5v6ZM3.975 5q-.425 0-.7-.288T3 4q0-.425.288-.713T4 3h16.025q.425 0 .7.288T21 4q0 .425-.288.713T20 5H3.975Zm0 16q-.425 0-.7-.288T3 20q0-.425.288-.713T4 19h16.025q.425 0 .7.288T21 20q0 .425-.288.713T20 21H3.975ZM5 9v6v-6Z"/>`),
		g.Group(children),
	)
}