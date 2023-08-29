package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHOneRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17q-.425 0-.713-.288T5 16V8q0-.425.288-.713T6 7q.425 0 .713.288T7 8v3h4V8q0-.425.288-.713T12 7q.425 0 .713.288T13 8v8q0 .425-.288.713T12 17q-.425 0-.713-.288T11 16v-3H7v3q0 .425-.288.713T6 17Zm12 0q-.425 0-.713-.288T17 16V9h-1q-.425 0-.713-.288T15 8q0-.425.288-.713T16 7h2q.425 0 .713.288T19 8v8q0 .425-.288.713T18 17Z"/>`),
		g.Group(children),
	)
}