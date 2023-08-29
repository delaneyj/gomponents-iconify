package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TamperDetectionOnOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 17v-2h2V3H4v3H2V1h16v6.5l4-4v11l-4-4V17h-4ZM3.625 20.025L-.25 16.15l1.4-1.4L3 16.6V8.5h1.5V13h1V7H7v6h1V8h1.5v5h1V9H12v11l-8.375.025ZM16 3v12V3Z"/>`),
		g.Group(children),
	)
}