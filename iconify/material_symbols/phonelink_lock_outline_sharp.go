package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonelinkLockOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 16v-5h1v-1q0-.825.588-1.413T18 8q.825 0 1.413.588T20 10v1h1v5h-6Zm2-5h2v-1q0-.425-.288-.713T18 9q-.425 0-.713.288T17 10v1ZM5 23V1h14v6h-2V6H7v12h10v-1h2v6H5Zm2-3v1h10v-1H7ZM7 4h10V3H7v1Zm0 0V3v1Zm0 16v1v-1Z"/>`),
		g.Group(children),
	)
}