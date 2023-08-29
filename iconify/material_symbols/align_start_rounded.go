package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignStartRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 4q-.425 0-.713-.288T2 3q0-.425.288-.713T3 2h18q.425 0 .713.288T22 3q0 .425-.288.713T21 4H3Zm5 6q-.425 0-.713-.288T7 9V8q0-.425.288-.713T8 7h8q.425 0 .713.288T17 8v1q0 .425-.288.713T16 10H8Zm0 6q-.425 0-.713-.288T7 15v-1q0-.425.288-.713T8 13h8q.425 0 .713.288T17 14v1q0 .425-.288.713T16 16H8Z"/>`),
		g.Group(children),
	)
}