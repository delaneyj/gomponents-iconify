package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubtitlesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm3-4h6q.425 0 .713-.288T14 15q0-.425-.288-.713T13 14H7q-.425 0-.713.288T6 15q0 .425.288.713T7 16Zm4-4h6q.425 0 .713-.288T18 11q0-.425-.288-.713T17 10h-6q-.425 0-.713.288T10 11q0 .425.288.713T11 12Zm-4 0q.425 0 .713-.288T8 11q0-.425-.288-.713T7 10q-.425 0-.713.288T6 11q0 .425.288.713T7 12Zm10 4q.425 0 .713-.288T18 15q0-.425-.288-.713T17 14q-.425 0-.713.288T16 15q0 .425.288.713T17 16Z"/>`),
		g.Group(children),
	)
}