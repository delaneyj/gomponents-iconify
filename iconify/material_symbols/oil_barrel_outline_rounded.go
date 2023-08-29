package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilBarrelOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h1v-6H4q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h1V5H4q-.425 0-.713-.288T3 4q0-.425.288-.713T4 3h16q.425 0 .713.288T21 4q0 .425-.288.713T20 5h-1v6h1q.425 0 .713.288T21 12q0 .425-.288.713T20 13h-1v6h1q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Zm3-2h10v-6q-.425 0-.713-.288T16 12q0-.425.288-.713T17 11V5H7v6q.425 0 .713.288T8 12q0 .425-.288.713T7 13v6Zm5-3q1.25 0 2.125-.863T15 13.05q0-.8-.338-1.35t-1.387-1.75l-.525-.575q-.3-.35-.75-.35t-.75.35l-.5.575q-1.05 1.2-1.4 1.75T9 13.05q0 1.225.875 2.087T12 16Zm-5 3V5v14Z"/>`),
		g.Group(children),
	)
}