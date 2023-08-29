package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncAltRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5.825 17l1.9 1.9q.3.3.288.7t-.313.7q-.3.275-.7.288t-.7-.288l-3.6-3.6q-.15-.15-.213-.325T2.426 16q0-.2.063-.375T2.7 15.3l3.6-3.6q.275-.275.688-.275t.712.275q.3.3.3.713t-.3.712L5.825 15H20q.425 0 .713.288T21 16q0 .425-.288.713T20 17H5.825Zm12.35-8H4q-.425 0-.713-.288T3 8q0-.425.288-.713T4 7h14.175l-1.9-1.9q-.3-.3-.287-.7t.312-.7q.3-.275.7-.288t.7.288l3.6 3.6q.15.15.213.325t.062.375q0 .2-.062.375T21.3 8.7l-3.6 3.6q-.275.275-.688.275T16.3 12.3q-.3-.3-.3-.713t.3-.712L18.175 9Z"/>`),
		g.Group(children),
	)
}