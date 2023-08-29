package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FoundationRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19v-2H4q-.425 0-.713-.288T3 16q0-.425.288-.713T4 15h1v-3H3.3q-.35 0-.475-.325t.15-.55l8.35-7.525q.275-.275.675-.275t.675.275l8.35 7.525q.275.225.15.55T20.7 12H19v3h1q.425 0 .712.288T21 16q0 .425-.288.713T20 17h-1v2q0 .425-.288.713T18 20q-.425 0-.713-.288T17 19v-2h-4v2q0 .425-.288.713T12 20q-.425 0-.713-.288T11 19v-2H7v2q0 .425-.288.713T6 20q-.425 0-.713-.288T5 19Zm2-4h4V6.6l-4 3.6V15Zm6 0h4v-4.8l-4-3.6V15Z"/>`),
		g.Group(children),
	)
}