package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SdOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-2h16V6H4v12Zm3-3h3q.425 0 .713-.288T11 14v-1.5q0-.425-.288-.713T10 11.5H7.5v-1h2v.5H11v-1q0-.425-.288-.713T10 9H7q-.425 0-.713.288T6 10v1.5q0 .425.288.713T7 12.5h2.5v1h-2V13H6v1q0 .425.288.713T7 15Zm6 0h4q.425 0 .713-.288T18 14v-4q0-.425-.288-.713T17 9h-4v6Zm1.5-1.5v-3h2v3h-2ZM4 18V6v12Z"/>`),
		g.Group(children),
	)
}