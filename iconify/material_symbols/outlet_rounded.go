package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutletRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 12q.425 0 .713-.288T10 11V8q0-.425-.288-.713T9 7q-.425 0-.713.288T8 8v3q0 .425.288.713T9 12Zm2 6h2q.425 0 .713-.288T14 17v-1q0-.825-.588-1.413T12 14q-.825 0-1.413.588T10 16v1q0 .425.288.713T11 18Zm4-6q.425 0 .713-.288T16 11V8q0-.425-.288-.713T15 7q-.425 0-.713.288T14 8v3q0 .425.288.713T15 12Zm-3 10q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}