package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteStackAddOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7 22l.025-15.025L22 7v10l-5 5H7Zm2-2h7l4-4V9H9v11Zm-4.7-.925L1.675 4.3L16.45 1.675L17.05 5H15l-.175-1L4 5.925l1 5.65v7.375l-.7.125Zm9.2-.575h2v-3h3v-2h-3v-3h-2v3h-3v2h3v3Zm1-4Z"/>`),
		g.Group(children),
	)
}