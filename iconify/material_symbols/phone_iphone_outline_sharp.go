package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhoneIphoneOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 23V1h14v22H5Zm2-5v3h10v-3H7Zm5 2.5q.425 0 .713-.288T13 19.5q0-.425-.288-.713T12 18.5q-.425 0-.713.288T11 19.5q0 .425.288.713T12 20.5ZM7 16h10V6H7v10ZM7 4h10V3H7v1Zm0 14v3v-3ZM7 4V3v1Z"/>`),
		g.Group(children),
	)
}