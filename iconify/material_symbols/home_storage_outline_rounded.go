package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeStorageOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.7 21q-.725 0-1.288-.475t-.687-1.2L3.2 10.175q-.075-.45.213-.812T4.175 9h15.65q.475 0 .763.363t.212.812l-1.525 9.15q-.125.725-.687 1.2T17.3 21H6.7ZM5.4 11l1.275 8h10.65l1.275-8H5.4Zm4.6 4h4q.425 0 .713-.288T15 14q0-.425-.288-.713T14 13h-4q-.425 0-.713.288T9 14q0 .425.288.713T10 15ZM6 8q-.425 0-.713-.288T5 7q0-.425.288-.713T6 6h12q.425 0 .713.288T19 7q0 .425-.288.713T18 8H6Zm2-3q-.425 0-.713-.288T7 4q0-.425.288-.713T8 3h8q.425 0 .713.288T17 4q0 .425-.288.713T16 5H8ZM6.675 19h10.65h-10.65Z"/>`),
		g.Group(children),
	)
}