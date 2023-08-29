package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorHeartOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 9q-.425 0-.713-.288T2 8V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v2q0 .425-.288.713T21 9q-.425 0-.713-.288T20 8V6H4v2q0 .425-.288.713T3 9Zm1 11q-.825 0-1.413-.588T2 18v-2q0-.425.288-.713T3 15q.425 0 .713.288T4 16v2h16v-2q0-.425.288-.713T21 15q.425 0 .713.288T22 16v2q0 .825-.588 1.413T20 20H4Zm6-3q.275 0 .525-.138t.375-.412l3.1-6.2l1.1 2.2q.125.275.375.413T16 13h5q.425 0 .713-.288T22 12q0-.425-.288-.713T21 11h-4.375L14.9 7.55q-.275-.5-.9-.5t-.9.5l-3.1 6.2l-1.1-2.2q-.125-.275-.375-.413T8 11H3q-.425 0-.713.288T2 12q0 .425.288.713T3 13h4.375L9.1 16.45q.125.275.375.413T10 17Zm2-5Z"/>`),
		g.Group(children),
	)
}