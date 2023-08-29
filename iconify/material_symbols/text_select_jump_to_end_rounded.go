package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectJumpToEndRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.175 13H4q-.425 0-.713-.288T3 12q0-.425.288-.713T4 11h6.175l-.9-.9Q9 9.825 9 9.412t.3-.712q.275-.275.7-.275t.7.275l2.6 2.6q.3.3.3.7t-.3.7l-2.6 2.6q-.275.275-.687.288T9.3 15.3q-.275-.275-.275-.7t.275-.7l.875-.9ZM19 20V4q0-.425.288-.713T20 3q.425 0 .713.288T21 4v16q0 .425-.288.713T20 21q-.425 0-.713-.288T19 20ZM15 5V3h2v2h-2Zm0 16v-2h2v2h-2ZM11 5V3h2v2h-2Zm0 16v-2h2v2h-2ZM7 5V3h2v2H7Zm0 16v-2h2v2H7Zm-2 0q-.825 0-1.413-.588T3 19h2v2ZM3 5q0-.825.588-1.413T5 3v2H3Z"/>`),
		g.Group(children),
	)
}