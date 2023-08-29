package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectedTvRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 16h2q0-.825-.588-1.413T5 14v2Zm3.55 0H10q0-2.075-1.463-3.538T5 11v1.45q1.475 0 2.513 1.038T8.55 16Zm3 0H13q0-1.65-.625-3.113t-1.713-2.55Q9.575 9.25 8.113 8.625T5 8v1.45q2.725 0 4.638 1.913T11.55 16ZM4 19q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v12q0 .825-.588 1.413T20 19h-4v1q0 .425-.288.713T15 21H9q-.425 0-.713-.288T8 20v-1H4Z"/>`),
		g.Group(children),
	)
}