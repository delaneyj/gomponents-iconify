package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHFourRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 17q-.425 0-.713-.288T2.5 16V8q0-.425.288-.713T3.5 7q.425 0 .713.288T4.5 8v3h4V8q0-.425.288-.713T9.5 7q.425 0 .713.288T10.5 8v8q0 .425-.288.713T9.5 17q-.425 0-.713-.288T8.5 16v-3h-4v3q0 .425-.288.713T3.5 17Zm15 0q-.425 0-.713-.288T17.5 16v-2h-4q-.425 0-.713-.288T12.5 13V8q0-.425.288-.713T13.5 7q.425 0 .713.288T14.5 8v4h3V8q0-.425.288-.713T18.5 7q.425 0 .713.288T19.5 8v4h1q.425 0 .713.288T21.5 13q0 .425-.288.713T20.5 14h-1v2q0 .425-.288.713T18.5 17Z"/>`),
		g.Group(children),
	)
}