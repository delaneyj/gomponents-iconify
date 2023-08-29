package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkLockedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 22v-5h1v-1q0-.825.588-1.413T20 14q.825 0 1.413.588T22 16v1h1v5h-6Zm2-5h2v-1q0-.425-.288-.713T20 15q-.425 0-.713.288T19 16v1ZM2 22L22 2v10h-2q-2.075 0-3.538 1.463T15 17v5H2Z"/>`),
		g.Group(children),
	)
}