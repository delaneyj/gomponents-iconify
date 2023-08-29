package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicationLiquid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 5V3h12v2H3Zm4.5 12.5h3V15H13v-3h-2.5V9.5h-3V12H5v3h2.5v2.5ZM4 21q-.825 0-1.413-.588T2 19V8q0-.825.588-1.413T4 6h10q.825 0 1.413.588T16 8v11q0 .825-.588 1.413T14 21H4Zm15-7.25q-.875-.425-1.438-1.413T17 10q0-1.7.863-2.85T20 6q1.275 0 2.138 1.15T23 10q0 1.35-.563 2.337T21 13.75V20q0 .425-.288.713T20 21q-.425 0-.713-.288T19 20v-6.25Z"/>`),
		g.Group(children),
	)
}