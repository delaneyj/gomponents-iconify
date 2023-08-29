package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PipRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 11q-.425 0-.713-.288T2 10q0-.425.288-.713T3 9h2.6L2 5.425q-.3-.3-.3-.712T2 4q.3-.3.713-.3t.712.3L7 7.6V5q0-.425.288-.713T8 4q.425 0 .713.288T9 5v5q0 .425-.288.713T8 11H3Zm1 9q-.825 0-1.413-.588T2 18v-4q0-.425.288-.713T3 13q.425 0 .713.288T4 14v4h7q.425 0 .713.288T12 19q0 .425-.288.713T11 20H4Zm17-7q-.425 0-.713-.288T20 12V6h-8q-.425 0-.713-.288T11 5q0-.425.288-.713T12 4h8q.825 0 1.413.588T22 6v6q0 .425-.288.713T21 13Zm-6 7q-.425 0-.713-.288T14 19v-3q0-.425.288-.713T15 15h6q.425 0 .713.288T22 16v3q0 .425-.288.713T21 20h-6Z"/>`),
		g.Group(children),
	)
}