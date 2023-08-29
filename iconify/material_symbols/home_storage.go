package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeStorage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21L3 9h18l-2 12H5Zm5-6h4q.425 0 .713-.288T15 14q0-.425-.288-.713T14 13h-4q-.425 0-.713.288T9 14q0 .425.288.713T10 15ZM6 8q-.425 0-.713-.288T5 7q0-.425.288-.713T6 6h12q.425 0 .713.288T19 7q0 .425-.288.713T18 8H6Zm2-3q-.425 0-.713-.288T7 4q0-.425.288-.713T8 3h8q.425 0 .713.288T17 4q0 .425-.288.713T16 5H8Z"/>`),
		g.Group(children),
	)
}