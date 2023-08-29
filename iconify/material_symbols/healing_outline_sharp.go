package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HealingOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.85 22.8L1.2 17.15L6.35 12L1.2 6.85L6.85 1.2L12 6.35l5.15-5.15l5.65 5.65L17.65 12l5.15 5.15l-5.65 5.65L12 17.65L6.85 22.8ZM12 11q.425 0 .713-.288T13 10q0-.425-.288-.713T12 9q-.425 0-.713.288T11 10q0 .425.288.713T12 11Zm-4.25-.4l2.85-2.85L6.85 4L4 6.85l3.75 3.75ZM10 13q.425 0 .713-.288T11 12q0-.425-.288-.713T10 11q-.425 0-.713.288T9 12q0 .425.288.713T10 13Zm2 2q.425 0 .713-.288T13 14q0-.425-.288-.713T12 13q-.425 0-.713.288T11 14q0 .425.288.713T12 15Zm2-2q.425 0 .713-.288T15 12q0-.425-.288-.713T14 11q-.425 0-.713.288T13 12q0 .425.288.713T14 13Zm3.15 7L20 17.15l-3.75-3.75l-2.85 2.85L17.15 20ZM8.475 8.45Zm7.05 7.075Z"/>`),
		g.Group(children),
	)
}