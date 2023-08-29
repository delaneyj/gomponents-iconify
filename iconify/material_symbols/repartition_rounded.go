package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RepartitionRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.825 8l.875.875q.3.3.3.713t-.3.712q-.3.3-.713.3t-.712-.3L2.7 7.7q-.3-.3-.3-.7t.3-.7l2.6-2.6q.275-.275.688-.275T6.7 3.7q.3.3.3.712t-.3.713L5.825 6H17q1.65 0 2.825 1.175T21 10q0 1.65-1.175 2.825T17 14H4q-.425 0-.713-.288T3 13q0-.425.288-.713T4 12h13q.825 0 1.413-.588T19 10q0-.825-.588-1.413T17 8H5.825ZM3 20v-2q0-.825.588-1.413T5 16h14q.825 0 1.413.588T21 18v2q0 .825-.588 1.413T19 22H5q-.825 0-1.413-.588T3 20Zm2 0h3.325v-2H5v2Zm5.325 0h3.325v-2h-3.325v2Zm5.35 0H19v-2h-3.325v2Z"/>`),
		g.Group(children),
	)
}