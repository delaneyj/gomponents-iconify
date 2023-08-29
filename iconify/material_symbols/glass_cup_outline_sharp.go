package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlassCupOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.2 22L3 2h18l-2.2 20H5.2Zm1.575-4L7 20h10l.225-2H6.775Zm-.2-2H17.45l1.3-12H5.25l1.325 12Zm.2 4h10.45h-10.45Z"/>`),
		g.Group(children),
	)
}