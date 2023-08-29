package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextWrapRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 20q-.425 0-.713-.288T4 19V5q0-.425.288-.713T5 4q.425 0 .713.288T6 5v14q0 .425-.288.713T5 20Zm14 0q-.425 0-.713-.288T18 19V5q0-.425.288-.713T19 4q.425 0 .713.288T20 5v14q0 .425-.288.713T19 20Zm-9.125-3.175L7.75 14.7q-.3-.3-.3-.7t.3-.7l2.125-2.125q.3-.3.713-.3t.712.3q.3.3.3.7t-.3.7l-.425.425H13q.825 0 1.413-.588T15 11q0-.825-.588-1.413T13 9H8q-.425 0-.713-.288T7 8q0-.425.288-.713T8 7h5q1.65 0 2.825 1.175T17 11q0 1.65-1.175 2.825T13 15h-2.125l.425.425q.3.3.3.7t-.3.7q-.3.3-.713.3t-.712-.3Z"/>`),
		g.Group(children),
	)
}