package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectAllRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19h2v2ZM3 5q0-.825.588-1.413T5 3v2H3Zm4 16v-2h2v2H7ZM7 5V3h2v2H7Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2q0 .825-.588 1.413T19 21Zm0-16V3q.825 0 1.413.588T21 5h-2ZM8 17q-.425 0-.713-.288T7 16V8q0-.425.288-.713T8 7h8q.425 0 .713.288T17 8v8q0 .425-.288.713T16 17H8Zm1-2h6V9H9v6Zm-6 2v-2h2v2H3Zm0-4v-2h2v2H3Zm0-4V7h2v2H3Zm16 8v-2h2v2h-2Zm0-4v-2h2v2h-2Zm0-4V7h2v2h-2Z"/>`),
		g.Group(children),
	)
}