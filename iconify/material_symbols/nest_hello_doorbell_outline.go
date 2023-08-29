package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestHelloDoorbellOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.538-1.463T7 17V7q0-2.075 1.463-3.538T12 2q2.075 0 3.538 1.463T17 7v10q0 2.075-1.463 3.538T12 22Zm0-2q1.25 0 2.125-.875T15 17V7q0-1.25-.875-2.125T12 4q-1.25 0-2.125.875T9 7v10q0 1.25.875 2.125T12 20Zm0-10q.825 0 1.413-.588T14 8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8q0 .825.588 1.413T12 10Zm0 2Zm0 6q.825 0 1.413-.588T14 16q0-.825-.588-1.413T12 14q-.825 0-1.413.588T10 16q0 .825.588 1.413T12 18Zm0-1q-.425 0-.713-.288T11 16q0-.425.288-.713T12 15q.425 0 .713.288T13 16q0 .425-.288.713T12 17Z"/>`),
		g.Group(children),
	)
}