package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkAsUnreadOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 17V6.25L10.5 2l9.8 5h-4.375L10.5 4.25L4 7.475V17H2Zm3 4V8h17v13H5Zm8.5-5.65L7 12v7h13v-7l-6.5 3.35Zm0-2L20 10H7l6.5 3.35ZM20 10H7h13Z"/>`),
		g.Group(children),
	)
}