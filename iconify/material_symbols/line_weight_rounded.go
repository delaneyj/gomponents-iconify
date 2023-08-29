package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineWeightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 20q-.2 0-.35-.15T3 19.5q0-.2.15-.35T3.5 19h17q.2 0 .35.15t.15.35q0 .2-.15.35t-.35.15h-17Zm.5-3q-.425 0-.713-.288T3 16q0-.425.288-.713T4 15h16q.425 0 .713.288T21 16q0 .425-.288.713T20 17H4Zm0-4q-.425 0-.713-.288T3 12v-1q0-.425.288-.713T4 10h16q.425 0 .713.288T21 11v1q0 .425-.288.713T20 13H4Zm0-5q-.425 0-.713-.288T3 7V5q0-.425.288-.713T4 4h16q.425 0 .713.288T21 5v2q0 .425-.288.713T20 8H4Z"/>`),
		g.Group(children),
	)
}