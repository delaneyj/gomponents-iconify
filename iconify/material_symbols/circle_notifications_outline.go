package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleNotificationsOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 18.5q.625 0 1.063-.438T13.5 17h-3q0 .625.438 1.063T12 18.5ZM7 16h10v-2h-1v-2.6q0-1.525-.788-2.788T13 7v-.5q0-.425-.288-.713T12 5.5q-.425 0-.713.288T11 6.5V7q-1.425.35-2.212 1.613T8 11.4V14H7v2Zm3-2v-3q0-.825.588-1.413T12 9q.825 0 1.413.588T14 11v3h-4Zm2 8q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Zm0-2q3.35 0 5.675-2.325T20 12q0-3.35-2.325-5.675T12 4Q8.65 4 6.325 6.325T4 12q0 3.35 2.325 5.675T12 20Zm0-8Z"/>`),
		g.Group(children),
	)
}