package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestDoorbellVisitor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 22q-2.075 0-3.538-1.463T7 17V7q0-2.075 1.463-3.538T12 2q2.075 0 3.538 1.463T17 7v5h-3q-.825 0-1.413.588T12 14q-.825 0-1.413.588T10 16q0 .8.588 1.375T12 18v4Zm2 0v-.575q0-1.3 1.25-1.875t2.75-.575q1.5 0 2.75.575T22 21.425V22h-8Zm4-4q-.825 0-1.413-.588T16 16q0-.825.588-1.413T18 14q.825 0 1.413.588T20 16q0 .825-.588 1.413T18 18Zm-6-1q-.425 0-.713-.288T11 16q0-.425.288-.713T12 15q.425 0 .713.288T13 16q0 .425-.288.713T12 17Zm0-7q.825 0 1.413-.588T14 8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8q0 .825.588 1.413T12 10Z"/>`),
		g.Group(children),
	)
}