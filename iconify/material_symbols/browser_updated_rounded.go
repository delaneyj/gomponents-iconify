package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrowserUpdatedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18q-.825 0-1.413-.588T2 16V5q0-.825.588-1.413T4 3h7q.425 0 .713.288T12 4q0 .425-.288.713T11 5H4v11h16v-2q0-.425.288-.713T21 13q.425 0 .713.288T22 14v2q0 .825-.588 1.413T20 18h-3l.7.7q.15.15.225.338t.075.387V20q0 .425-.288.712T17 21H7q-.425 0-.713-.288T6 20v-.575q0-.2.075-.388T6.3 18.7L7 18H4Zm10-6.825V4q0-.425.288-.713T15 3q.425 0 .713.288T16 4v7.175L17.9 9.3q.275-.275.688-.288t.712.288q.275.275.275.7t-.275.7L15 15l-4.3-4.3q-.275-.275-.287-.687T10.7 9.3q.275-.275.7-.275t.7.275l1.9 1.875Z"/>`),
		g.Group(children),
	)
}