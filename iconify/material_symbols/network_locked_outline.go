package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkLockedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22L22 2v10h-2V6.825L6.825 20H15v2H2Zm15.85 0q-.35 0-.6-.25t-.25-.6v-3.3q0-.35.25-.6t.6-.25H18v-1q0-.825.588-1.413T20 14q.825 0 1.413.588T22 16v1h.15q.35 0 .6.25t.25.6v3.3q0 .35-.25.6t-.6.25h-4.3ZM19 17h2v-1q0-.425-.288-.713T20 15q-.425 0-.713.288T19 16v1Zm-5.575-3.575Z"/>`),
		g.Group(children),
	)
}