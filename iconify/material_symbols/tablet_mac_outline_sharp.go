package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TabletMacOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 23V1h18v22H3Zm2-5v3h14v-3H5Zm0-2h14V6H5v10ZM5 4h14V3H5v1Zm0 0V3v1Zm0 14v3v-3Zm7 2.5q.425 0 .713-.288T13 19.5q0-.425-.288-.713T12 18.5q-.425 0-.713.288T11 19.5q0 .425.288.713T12 20.5Z"/>`),
		g.Group(children),
	)
}