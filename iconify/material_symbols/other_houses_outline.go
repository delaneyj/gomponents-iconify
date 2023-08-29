package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OtherHousesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21v-9.375L2.2 13L1 11.4L12 3l11 8.4l-1.2 1.575l-1.8-1.35V21H4Zm2-2h12v-8.9l-6-4.575L6 10.1V19Zm0 0h12H6Zm2-4q-.425 0-.713-.288T7 14q0-.425.288-.713T8 13q.425 0 .713.288T9 14q0 .425-.288.713T8 15Zm4 0q-.425 0-.713-.288T11 14q0-.425.288-.713T12 13q.425 0 .713.288T13 14q0 .425-.288.713T12 15Zm4 0q-.425 0-.713-.288T15 14q0-.425.288-.713T16 13q.425 0 .713.288T17 14q0 .425-.288.713T16 15Z"/>`),
		g.Group(children),
	)
}