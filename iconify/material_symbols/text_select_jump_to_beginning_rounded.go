package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectJumpToBeginningRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13.3 15.3l-2.6-2.6q-.3-.3-.3-.7t.3-.7l2.6-2.6q.275-.275.688-.287t.712.287q.275.275.275.7t-.275.7l-.875.9H20q.425 0 .713.288T21 12q0 .425-.288.713T20 13h-6.175l.9.9q.275.275.275.688t-.3.712q-.275.275-.7.275t-.7-.275ZM4 21q-.425 0-.713-.288T3 20V4q0-.425.288-.713T4 3q.425 0 .713.288T5 4v16q0 .425-.288.713T4 21Zm3 0v-2h2v2H7ZM7 5V3h2v2H7Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 0V3q.825 0 1.413.588T21 5h-2Zm0 16v-2h2q0 .825-.588 1.413T19 21Z"/>`),
		g.Group(children),
	)
}