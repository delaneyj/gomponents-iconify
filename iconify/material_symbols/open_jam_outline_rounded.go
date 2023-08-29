package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenJamOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 19H7q-.425 0-.713.288T6 20q0 .425.288.713T7 21h10q.425 0 .713-.288T18 20q0-.425-.288-.713T17 19h-4v-7.2l.9.9q.275.275.7.275t.7-.275q.275-.275.275-.7t-.275-.7l-2.6-2.6q-.3-.3-.7-.3t-.7.3l-2.6 2.6q-.275.275-.275.7t.275.7q.275.275.7.275t.7-.275l.9-.9V19Zm1-7Zm-8 4q-.825 0-1.413-.588T2 14V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v9q0 .825-.588 1.413T20 16h-4q-.425 0-.713-.288T15 15q0-.425.288-.713T16 14h4V5H4v9h4q.425 0 .713.288T9 15q0 .425-.288.713T8 16H4Z"/>`),
		g.Group(children),
	)
}