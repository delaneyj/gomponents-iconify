package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProcedureOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 11q-.95 0-2.025-.537t-2-1.488q-.925-.95-1.462-2.025t-.538-2q0-.425.125-.775t.375-.6q.65-.65 2.638-1.138t3.737-.412q.625.05 1.038.15t.612.3q.175.175.288.537t.162.913q.125 1.725-.35 3.8t-1.15 2.75q-.25.25-.637.388T19 11Zm.875-5q.05-.45.075-.95T20 4q-.5-.025-1.012.013t-1.013.087q.275.2.537.425t.488.45q.25.25.463.5t.412.525ZM19 9q.025-.5-.375-1.238T17.575 6.4q-.625-.625-1.338-1.012T15.026 5q.05.575.425 1.3t.925 1.275q.6.6 1.313.988T19 9Zm2.3 9.7L16.6 14h-10L1.3 8.7l1.4-1.4L7.4 12h10l5.3 5.3l-1.4 1.4ZM8 22v-4q0-.825.588-1.413T10 16h4q.825 0 1.413.588T16 18v4H8Zm2-2h4v-2h-4v2Zm0 0v-2v2Z"/>`),
		g.Group(children),
	)
}