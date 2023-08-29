package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkAsUnreadSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 17V6.25L10.5 2l9.8 5h-4.375L10.5 4.25L4 7.475V17H2Zm3 4V8h17v13H5Zm8.5-5.65L20 12v-2l-6.5 3.35L7 10v2l6.5 3.35Z"/>`),
		g.Group(children),
	)
}