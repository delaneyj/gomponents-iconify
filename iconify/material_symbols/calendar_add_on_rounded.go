package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarAddOnRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 19h-2q-.425 0-.713-.288T14 18q0-.425.288-.713T15 17h2v-2q0-.425.288-.713T18 14q.425 0 .713.288T19 15v2h2q.425 0 .713.288T22 18q0 .425-.288.713T21 19h-2v2q0 .425-.288.713T18 22q-.425 0-.713-.288T17 21v-2ZM5 20q-.825 0-1.413-.588T3 18V6q0-.825.588-1.413T5 4h1V3q0-.425.288-.713T7 2q.425 0 .713.288T8 3v1h6V3q0-.425.288-.713T15 2q.425 0 .713.288T16 3v1h1q.825 0 1.413.588T19 6v6.1q-.5-.075-1-.075t-1 .075V10H5v8h7q0 .5.075 1t.275 1H5Z"/>`),
		g.Group(children),
	)
}