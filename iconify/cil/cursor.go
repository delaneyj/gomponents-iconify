package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cursor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M446.754 47.9a20.075 20.075 0 0 0-21.307-2.745L32 229.835v37l165.349 66.139L303.317 496h37L453.281 68.369a20.072 20.072 0 0 0-6.527-20.469Zm-129.63 410.624l-98.473-151.5l-148.9-59.561L415.779 85.044Z"/>`),
		g.Group(children),
	)
}