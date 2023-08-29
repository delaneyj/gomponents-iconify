package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataArrayRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 20q-.425 0-.713-.288T15 19q0-.425.288-.713T16 18h2V6h-2q-.425 0-.713-.288T15 5q0-.425.288-.713T16 4h2q.825 0 1.413.588T20 6v12q0 .825-.588 1.413T18 20h-2ZM6 20q-.825 0-1.413-.588T4 18V6q0-.825.588-1.413T6 4h2q.425 0 .713.288T9 5q0 .425-.288.713T8 6H6v12h2q.425 0 .713.288T9 19q0 .425-.288.713T8 20H6Z"/>`),
		g.Group(children),
	)
}