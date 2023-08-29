package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SosRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 17q-.825 0-1.413-.588T8.5 15V9q0-.825.588-1.413T10.5 7h3q.825 0 1.413.588T15.5 9v6q0 .825-.588 1.413T13.5 17h-3ZM5 17H2q-.425 0-.713-.288T1 16q0-.425.288-.713T2 15h3v-2H3q-.825 0-1.413-.588T1 11V9q0-.825.588-1.413T3 7h3q.425 0 .713.288T7 8q0 .425-.288.713T6 9H3v2h2q.825 0 1.413.588T7 13v2q0 .825-.588 1.413T5 17Zm16 0h-3q-.425 0-.713-.288T17 16q0-.425.288-.713T18 15h3v-2h-2q-.825 0-1.413-.588T17 11V9q0-.825.588-1.413T19 7h3q.425 0 .713.288T23 8q0 .425-.288.713T22 9h-3v2h2q.825 0 1.413.588T23 13v2q0 .825-.588 1.413T21 17Zm-10.5-2h3V9h-3v6Z"/>`),
		g.Group(children),
	)
}