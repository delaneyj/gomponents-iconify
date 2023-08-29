package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrowaveGen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm1-3h10V7H5v10Zm13 0q.425 0 .713-.288T19 16q0-.425-.288-.713T18 15q-.425 0-.713.288T17 16q0 .425.288.713T18 17ZM7 15V9h6v6H7Zm11-2q.425 0 .713-.288T19 12q0-.425-.288-.713T18 11q-.425 0-.713.288T17 12q0 .425.288.713T18 13Zm0-4q.425 0 .713-.288T19 8q0-.425-.288-.713T18 7q-.425 0-.713.288T17 8q0 .425.288.713T18 9Z"/>`),
		g.Group(children),
	)
}