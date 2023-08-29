package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealingSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.85 22.8L1.2 17.15L17.15 1.2l5.65 5.65L6.85 22.8Zm10.3 0l-4.45-4.45l5.65-5.65l4.45 4.45l-5.65 5.65ZM12 15q.425 0 .713-.288T13 14q0-.425-.288-.713T12 13q-.425 0-.713.288T11 14q0 .425.288.713T12 15Zm-2-2q.425 0 .713-.288T11 12q0-.425-.288-.713T10 11q-.425 0-.713.288T9 12q0 .425.288.713T10 13Zm4 0q.425 0 .713-.288T15 12q0-.425-.288-.713T14 11q-.425 0-.713.288T13 12q0 .425.288.713T14 13Zm-8.375-1.725L1.2 6.85L6.85 1.2l4.425 4.425l-5.65 5.65ZM12 11q.425 0 .713-.288T13 10q0-.425-.288-.713T12 9q-.425 0-.713.288T11 10q0 .425.288.713T12 11Z"/>`),
		g.Group(children),
	)
}