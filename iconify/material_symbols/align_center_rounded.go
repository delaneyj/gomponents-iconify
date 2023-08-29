package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 13q-.425 0-.713-.288T2 12q0-.425.288-.713T3 11h18q.425 0 .713.288T22 12q0 .425-.288.713T21 13H3Zm5-3q-.425 0-.713-.288T7 9V8q0-.425.288-.713T8 7h8q.425 0 .713.288T17 8v1q0 .425-.288.713T16 10H8Zm0 7q-.425 0-.713-.288T7 16v-1q0-.425.288-.713T8 14h8q.425 0 .713.288T17 15v1q0 .425-.288.713T16 17H8Z"/>`),
		g.Group(children),
	)
}