package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenInFullRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.425 0-.713-.288T3 20v-6q0-.425.288-.713T4 13q.425 0 .713.288T5 14v3.6L17.6 5H14q-.425 0-.713-.288T13 4q0-.425.288-.713T14 3h6q.425 0 .713.288T21 4v6q0 .425-.288.713T20 11q-.425 0-.713-.288T19 10V6.4L6.4 19H10q.425 0 .713.288T11 20q0 .425-.288.713T10 21H4Z"/>`),
		g.Group(children),
	)
}