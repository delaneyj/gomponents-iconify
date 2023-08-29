package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreakfastDiningSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21V10.45q-.925-.55-1.463-1.462T2 7q0-1.65 1.175-2.825T6 3h12q1.65 0 2.825 1.175T22 7q0 1.075-.537 1.988T20 10.45V21H4Zm8-3.6l4.4-4.4L12 8.6L7.6 13l4.4 4.4Zm0-2.8L10.4 13l1.6-1.6l1.6 1.6l-1.6 1.6Z"/>`),
		g.Group(children),
	)
}