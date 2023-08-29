package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClosedCaptionOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20q-.825 0-1.413-.588T3 18V6q0-.825.588-1.413T5 4h14q.825 0 1.413.588T21 6v12q0 .825-.588 1.413T19 20H5Zm0-2h14V6H5v12Zm2-3h3q.425 0 .713-.288T11 14v-1H9.5v.5h-2v-3h2v.5H11v-1q0-.425-.288-.713T10 9H7q-.425 0-.713.288T6 10v4q0 .425.288.713T7 15Zm7 0h3q.425 0 .713-.288T18 14v-1h-1.5v.5h-2v-3h2v.5H18v-1q0-.425-.288-.713T17 9h-3q-.425 0-.713.288T13 10v4q0 .425.288.713T14 15Zm-9 3V6v12Z"/>`),
		g.Group(children),
	)
}