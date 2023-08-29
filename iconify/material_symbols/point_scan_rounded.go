package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PointScanRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14q-.825 0-1.413-.588T10 12q0-.825.588-1.413T12 10q.825 0 1.413.588T14 12q0 .825-.588 1.413T12 14Zm0-6q-.425 0-.713-.288T11 7V4q0-.425.288-.713T12 3q.425 0 .713.288T13 4v3q0 .425-.288.713T12 8Zm0 13q-.425 0-.713-.288T11 20v-3q0-.425.288-.713T12 16q.425 0 .713.288T13 17v3q0 .425-.288.713T12 21Zm5-8q-.425 0-.713-.288T16 12q0-.425.288-.713T17 11h3q.425 0 .713.288T21 12q0 .425-.288.713T20 13h-3ZM4 13q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h3q.425 0 .713.288T8 12q0 .425-.288.713T7 13H4Z"/>`),
		g.Group(children),
	)
}