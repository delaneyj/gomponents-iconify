package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourGMobiledataRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17q-.425 0-.713-.288T7 16v-2H4q-.425 0-.713-.288T3 13V8q0-.425.288-.713T4 7q.425 0 .713.288T5 8v4h2V8q0-.425.288-.713T8 7q.425 0 .713.288T9 8v4h1q.425 0 .713.288T11 13q0 .425-.288.713T10 14H9v2q0 .425-.288.713T8 17Zm6 0q-.825 0-1.413-.588T12 15V9q0-.825.588-1.413T14 7h6q.425 0 .713.288T21 8q0 .425-.288.713T20 9h-6v6h5v-2h-1.5q-.425 0-.713-.288T16.5 12q0-.425.288-.713T17.5 11H20q.425 0 .713.288T21 12v3q0 .825-.588 1.413T19 17h-5Z"/>`),
		g.Group(children),
	)
}