package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PowerInputRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 15q-.425 0-.713-.288T2 14q0-.425.288-.713T3 13h3q.425 0 .713.288T7 14q0 .425-.288.713T6 15H3Zm7 0q-.425 0-.713-.288T9 14q0-.425.288-.713T10 13h3q.425 0 .713.288T14 14q0 .425-.288.713T13 15h-3Zm7 0q-.425 0-.713-.288T16 14q0-.425.288-.713T17 13h3q.425 0 .713.288T21 14q0 .425-.288.713T20 15h-3ZM3 11q-.425 0-.713-.288T2 10q0-.425.288-.713T3 9h17q.425 0 .713.288T21 10q0 .425-.288.713T20 11H3Z"/>`),
		g.Group(children),
	)
}