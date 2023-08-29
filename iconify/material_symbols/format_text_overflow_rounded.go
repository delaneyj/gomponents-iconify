package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextOverflowRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20q-.425 0-.713-.288T4 19V5q0-.425.288-.713T5 4q.425 0 .713.288T6 5v14q0 .425-.288.713T5 20Zm8-11q-.425 0-.713-.288T12 8V5q0-.425.288-.713T13 4q.425 0 .713.288T14 5v3q0 .425-.288.713T13 9Zm4.175 5.825q-.3.3-.713.3t-.712-.3q-.3-.3-.3-.7t.3-.7l.425-.425H9q-.425 0-.713-.288T8 12q0-.425.288-.713T9 11h7.175l-.425-.425q-.3-.3-.3-.7t.3-.7q.3-.3.713-.3t.712.3L19.3 11.3q.3.3.3.7t-.3.7l-2.125 2.125ZM13 20q-.425 0-.713-.288T12 19v-3q0-.425.288-.713T13 15q.425 0 .713.288T14 16v3q0 .425-.288.713T13 20Z"/>`),
		g.Group(children),
	)
}