package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm3-5h3q.425 0 .713-.288T11 14v-1.5q0-.425-.288-.713T10 11.5H7.5v-1h2q0 .2.15.35T10 11h.5q.2 0 .35-.15t.15-.35V10q0-.425-.288-.713T10 9H7q-.425 0-.713.288T6 10v1.5q0 .425.288.713T7 12.5h2.5v1h-2q0-.2-.15-.35T7 13h-.5q-.2 0-.35.15T6 13.5v.5q0 .425.288.713T7 15Zm6.5 0H17q.425 0 .713-.288T18 14v-4q0-.425-.288-.713T17 9h-3.5q-.2 0-.35.15T13 9.5v5q0 .2.15.35t.35.15Zm1-1.5v-3h2v3h-2Z"/>`),
		g.Group(children),
	)
}