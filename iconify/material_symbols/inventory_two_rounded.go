package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InventoryTwoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20V8.725q-.45-.275-.725-.713T2 7V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v3q0 .575-.275 1.012T21 8.725V20q0 .825-.588 1.413T19 22H5ZM4 7h16V4H4v3Zm6 7h4q.425 0 .713-.288T15 13q0-.425-.288-.713T14 12h-4q-.425 0-.713.288T9 13q0 .425.288.713T10 14Z"/>`),
		g.Group(children),
	)
}