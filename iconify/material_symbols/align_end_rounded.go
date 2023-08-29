package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignEndRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21q0-.425.288-.713T3 20h18q.425 0 .713.288T22 21q0 .425-.288.713T21 22H3Zm5-11q-.425 0-.713-.288T7 10V9q0-.425.288-.713T8 8h8q.425 0 .713.288T17 9v1q0 .425-.288.713T16 11H8Zm0 6q-.425 0-.713-.288T7 16v-1q0-.425.288-.713T8 14h8q.425 0 .713.288T17 15v1q0 .425-.288.713T16 17H8Z"/>`),
		g.Group(children),
	)
}