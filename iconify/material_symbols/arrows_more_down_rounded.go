package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsMoreDownRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-.425 0-.713-.288T5 20v-9q0-.425.288-.713T6 10q.425 0 .713.288T7 11v8h8q.425 0 .713.288T16 20q0 .425-.288.713T15 21H6Zm5-5q-.425 0-.713-.288T10 15V6q0-.425.288-.713T11 5q.425 0 .713.288T12 6v8h8q.425 0 .713.288T21 15q0 .425-.288.713T20 16h-9Z"/>`),
		g.Group(children),
	)
}