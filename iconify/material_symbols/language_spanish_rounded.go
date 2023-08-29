package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LanguageSpanishRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 15h4q.425 0 .713.288T11 16q0 .425-.288.713T10 17H5q-.425 0-.713-.288T4 16V8q0-.425.288-.713T5 7h5q.425 0 .713.288T11 8q0 .425-.288.713T10 9H6v2h3q.425 0 .713.288T10 12q0 .425-.288.713T9 13H6v2Zm9 2q-.825 0-1.413-.588T13 15q0-.425.288-.713T14 14q.425 0 .713.288T15 15h3v-2h-3q-.825 0-1.413-.588T13 11V9q0-.825.588-1.413T15 7h3q.825 0 1.413.588T20 9q0 .425-.288.713T19 10q-.425 0-.713-.288T18 9h-3v2h3q.825 0 1.413.588T20 13v2q0 .825-.588 1.413T18 17h-3Z"/>`),
		g.Group(children),
	)
}