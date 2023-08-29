package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SweepOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 18q-.425 0-.713-.288T10 17q0-.425.288-.713T11 16h4q.425 0 .713.288T16 17q0 .425-.288.713T15 18h-4Zm-4.95-2.85L14.5 6.7q.3-.3.7-.3t.7.3q.3.3.3.712t-.3.713L6.75 17.3q-.3.3-.7.3t-.7-.3l-4.275-4.275q-.3-.3-.287-.713T1.1 11.6q.3-.3.713-.3t.712.3l3.525 3.55ZM15 14q-.425 0-.712-.288T14 13q0-.425.288-.713T15 12h4q.425 0 .713.288T20 13q0 .425-.288.713T19 14h-4Zm4-4q-.425 0-.713-.288T18 9q0-.425.288-.713T19 8h4q.425 0 .713.288T24 9q0 .425-.288.713T23 10h-4Z"/>`),
		g.Group(children),
	)
}