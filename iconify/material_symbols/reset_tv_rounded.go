package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ResetTvRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20v-1H4q-.825 0-1.413-.588T2 17V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v5h-9.2l1.15-1.15q.275-.275.275-.7t-.275-.7q-.275-.275-.7-.275t-.7.275L9.7 10.3q-.3.3-.3.7t.3.7l2.85 2.85q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7L12.8 12H22v5q0 .825-.588 1.413T20 19h-4v1q0 .425-.288.713T15 21H9q-.425 0-.713-.288T8 20Z"/>`),
		g.Group(children),
	)
}