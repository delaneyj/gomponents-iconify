package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HPlusMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 13v3q0 .425-.288.713T5 17q-.425 0-.713-.288T4 16V8q0-.425.288-.713T5 7q.425 0 .713.288T6 8v3h6V8q0-.425.288-.713T13 7q.425 0 .713.288T14 8v8q0 .425-.288.713T13 17q-.425 0-.713-.288T12 16v-3H6Zm12 0h-1q-.425 0-.713-.288T16 12q0-.425.288-.713T17 11h1v-1q0-.425.288-.713T19 9q.425 0 .713.288T20 10v1h1q.425 0 .713.288T22 12q0 .425-.288.713T21 13h-1v1q0 .425-.288.713T19 15q-.425 0-.713-.288T18 14v-1Z"/>`),
		g.Group(children),
	)
}