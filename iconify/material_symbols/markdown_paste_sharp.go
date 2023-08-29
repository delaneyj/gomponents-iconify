package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarkdownPasteSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21v-9h10v9h-2v-7h-2v5h-2v-5h-2v7h-2Zm-9 0V3h6.175q.275-.875 1.075-1.438T12 1q1 0 1.788.563T14.85 3H21v7h-2V5h-2v3H7V5H5v14h5v2H3Zm9-16q.425 0 .713-.288T13 4q0-.425-.288-.713T12 3q-.425 0-.713.288T11 4q0 .425.288.713T12 5Z"/>`),
		g.Group(children),
	)
}