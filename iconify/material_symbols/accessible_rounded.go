package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AccessibleRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Zm6 16q-.425 0-.713-.288T17 21v-4h-5q-.825 0-1.413-.588T10 15V9q0-.875.563-1.438T12 7q.6 0 1.038.263t.937.837q1.15 1.4 2.1 2.025t2.05.8q.375.05.625.325t.25.65q0 .475-.325.775t-.75.225q-1.075-.2-2.112-.713T14 11.05v3.45h3q.825 0 1.413.588T19 16.5V21q0 .425-.288.713T18 22Zm-8 0q-2.075 0-3.538-1.463T5 17q0-1.8 1.137-3.175T9 12.1v2.05q-.875.35-1.438 1.113T7 17q0 1.25.875 2.125T10 20q.975 0 1.738-.563T12.85 18h2.05q-.35 1.725-1.725 2.863T10 22Z"/>`),
		g.Group(children),
	)
}