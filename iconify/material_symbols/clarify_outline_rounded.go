package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClarifyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17h5q.425 0 .713-.288T13 16q0-.425-.288-.713T12 15H7q-.425 0-.713.288T6 16q0 .425.288.713T7 17ZM17 7q-.425 0-.713.288T16 8v8q0 .425.288.713T17 17q.425 0 .713-.288T18 16V8q0-.425-.288-.713T17 7ZM7 13h5q.425 0 .713-.288T13 12q0-.425-.288-.713T12 11H7q-.425 0-.713.288T6 12q0 .425.288.713T7 13Zm0-4h5q.425 0 .713-.288T13 8q0-.425-.288-.713T12 7H7q-.425 0-.713.288T6 8q0 .425.288.713T7 9ZM4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v14q0 .825-.588 1.413T20 21H4Zm0-2h16V5H4v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}