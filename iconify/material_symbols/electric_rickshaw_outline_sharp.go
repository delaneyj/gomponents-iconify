package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricRickshawOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17q-.975 0-1.738-.563T3.2 15H1V3h15.025L21 9v2.15q.875.3 1.438 1.088T23 14q0 1.25-.875 2.125T20 17q-.975 0-1.763-.563T17.15 15h-8.3q-.35.875-1.113 1.438T6 17ZM3 8h4V5H3v3Zm6 5h5V5H9v3h3v2H9v3Zm7-4h2.4L16 6.1V9ZM6 15q.425 0 .713-.288T7 14q0-.425-.288-.713T6 13q-.425 0-.713.288T5 14q0 .425.288.713T6 15Zm14 0q.425 0 .713-.288T21 14q0-.425-.288-.713T20 13q-.425 0-.713.288T19 14q0 .425.288.713T20 15Zm-7 8l-6-3h4v-2l6 3h-4v2ZM3 10v3h.15q.35-.875 1.113-1.438T6 11q.275 0 .525.038T7 11.15V10H3Zm13 3h1.15q.225-.65.713-1.137T19 11.15V11h-3v2ZM3 10h4h-4Zm13 1h3h-3Z"/>`),
		g.Group(children),
	)
}