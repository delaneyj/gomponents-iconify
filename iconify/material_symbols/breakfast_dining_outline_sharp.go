package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreakfastDiningOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V10.45q-.925-.55-1.463-1.462T2 7q0-1.65 1.175-2.825T6 3h12q1.65 0 2.825 1.175T22 7q0 1.075-.537 1.988T20 10.45V21H4Zm2-2h12V9.3l1-.6q.45-.275.725-.725T20 7q0-.825-.588-1.412T18 5H6q-.825 0-1.413.588T4 7q0 .55.263 1.012T5 8.75l1 .55V19Zm6-1.6l4.4-4.4L12 8.6L7.6 13l4.4 4.4Zm0-2.8L10.4 13l1.6-1.6l1.6 1.6l-1.6 1.6Zm0-2.6Z"/>`),
		g.Group(children),
	)
}