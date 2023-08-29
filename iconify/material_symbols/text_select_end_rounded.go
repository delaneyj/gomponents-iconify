package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectEndRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 5V3h2v2h-2Zm0 16v-2h2v2h-2ZM7 5V3h2v2H7Zm0 16v-2h2v2H7ZM3 9V7h2v2H3Zm0 4v-2h2v2H3Zm0 4v-2h2v2H3Zm13 4q-.425 0-.713-.288T15 20q0-.425.288-.713T16 19h1V5h-1q-.425 0-.713-.288T15 4q0-.425.288-.713T16 3h4q.425 0 .713.288T21 4q0 .425-.288.713T20 5h-1v14h1q.425 0 .713.288T21 20q0 .425-.288.713T20 21h-4ZM5 21q-.825 0-1.413-.588T3 19h2v2ZM3 5q0-.825.588-1.413T5 3v2H3Z"/>`),
		g.Group(children),
	)
}