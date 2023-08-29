package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardwareSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 11l-3-3v3H9V8H4q0-2.075 1.463-3.538T9 3h6v3l3-3h2v8h-2ZM9 21v-8h6v8H9Z"/>`),
		g.Group(children),
	)
}