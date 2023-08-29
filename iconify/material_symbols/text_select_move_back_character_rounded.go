package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectMoveBackCharacterRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 5V3h2v2h-2Zm0 16v-2h2v2h-2ZM7 5V3h2v2H7Zm0 16v-2h2v2H7Zm-.7-5.7l-2.6-2.6q-.3-.3-.3-.7t.3-.7l2.6-2.6q.275-.275.688-.287T7.7 8.7q.275.275.275.7t-.275.7l-.875.9H13q.425 0 .713.288T14 12q0 .425-.288.713T13 13H6.825l.9.9q.275.275.275.688t-.3.712q-.275.275-.7.275t-.7-.275ZM16 21q-.425 0-.713-.288T15 20q0-.425.288-.713T16 19h1V5h-1q-.425 0-.713-.288T15 4q0-.425.288-.713T16 3h4q.425 0 .713.288T21 4q0 .425-.288.713T20 5h-1v14h1q.425 0 .713.288T21 20q0 .425-.288.713T20 21h-4ZM3 5q0-.825.588-1.413T5 3v2H3Zm2 16q-.825 0-1.413-.588T3 19h2v2Z"/>`),
		g.Group(children),
	)
}