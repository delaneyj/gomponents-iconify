package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CorporateFareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h6q.825 0 1.413.588T12 5v2h8q.825 0 1.413.588T22 9v10q0 .825-.588 1.413T20 21H4Zm0-2h6v-2H4v2Zm0-4h6v-2H4v2Zm0-4h6V9H4v2Zm0-4h6V5H4v2Zm8 12h8V9h-8v10Zm3-6q-.425 0-.713-.288T14 12q0-.425.288-.713T15 11h2q.425 0 .713.288T18 12q0 .425-.288.713T17 13h-2Zm0 4q-.425 0-.713-.288T14 16q0-.425.288-.713T15 15h2q.425 0 .713.288T18 16q0 .425-.288.713T17 17h-2Z"/>`),
		g.Group(children),
	)
}