package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusinessCenterOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V6h6V2h8v4h6v15H2Zm8-15h4V4h-4v2Zm10 9h-5v2H9v-2H4v4h16v-4Zm-9 0h2v-2h-2v2Zm-7-2h5v-2h6v2h5V8H4v5Zm8 1Z"/>`),
		g.Group(children),
	)
}