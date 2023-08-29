package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PictureAsPdfRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 10.5h1q.425 0 .713-.288T12 9.5v-1q0-.425-.288-.713T11 7.5H9.5q-.2 0-.35.15T9 8v4q0 .2.15.35t.35.15q.2 0 .35-.15T10 12v-1.5Zm0-1v-1h1v1h-1Zm5 3q.425 0 .713-.288T16 11.5v-3q0-.425-.288-.713T15 7.5h-1.5q-.2 0-.35.15T13 8v4q0 .2.15.35t.35.15H15Zm-1-1v-3h1v3h-1Zm4-1h.5q.2 0 .35-.15T19 10q0-.2-.15-.35t-.35-.15H18v-1h.5q.2 0 .35-.15T19 8q0-.2-.15-.35t-.35-.15h-1q-.2 0-.35.15T17 8v4q0 .2.15.35t.35.15q.2 0 .35-.15T18 12v-1.5ZM8 18q-.825 0-1.413-.588T6 16V4q0-.825.588-1.413T8 2h12q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H8Zm-4 4q-.825 0-1.413-.588T2 20V7q0-.425.288-.713T3 6q.425 0 .713.288T4 7v13h13q.425 0 .713.288T18 21q0 .425-.288.713T17 22H4Z"/>`),
		g.Group(children),
	)
}