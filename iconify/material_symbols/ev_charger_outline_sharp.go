package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EvChargerOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.5 19l2.5-4H9.5v-3L7 16h1.5v3ZM6 10h6V5H6v5Zm0 9h6v-7H6v7Zm-2 2V3h10v9h3v7.5h2.25V9H18V6h.5V4.5h1V6h1V4.5h1V6h.5v3h-1.25v12H15.5v-7.5H14V21H4Zm8-2H6h6Z"/>`),
		g.Group(children),
	)
}