package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneIphoneSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 23V1h14v22H5Zm7-2.5q.425 0 .713-.288T13 19.5q0-.425-.288-.713T12 18.5q-.425 0-.713.288T11 19.5q0 .425.288.713T12 20.5ZM7 16h10V6H7v10Z"/>`),
		g.Group(children),
	)
}