package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FactCheckOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v14q0 .825-.588 1.413T20 21H4Zm0-2h16V5H4v14Zm5-2q.425 0 .713-.288T10 16q0-.425-.288-.713T9 15H6q-.425 0-.713.288T5 16q0 .425.288.713T6 17h3Zm5.55-4.825l-.725-.725q-.3-.3-.7-.287t-.7.312q-.275.3-.288.7t.288.7L13.85 14.3q.3.3.7.3t.7-.3l3.55-3.55q.3-.3.3-.7t-.3-.7q-.3-.3-.713-.3t-.712.3l-2.825 2.825ZM9 13q.425 0 .713-.288T10 12q0-.425-.288-.713T9 11H6q-.425 0-.713.288T5 12q0 .425.288.713T6 13h3Zm0-4q.425 0 .713-.288T10 8q0-.425-.288-.713T9 7H6q-.425 0-.713.288T5 8q0 .425.288.713T6 9h3ZM4 19V5v14Z"/>`),
		g.Group(children),
	)
}