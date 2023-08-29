package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InventoryTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20V8.725q-.45-.275-.725-.713T2 7V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v3q0 .575-.275 1.012T21 8.725V20q0 .825-.588 1.413T19 22H5ZM4 7h16V4H4v3Zm5 7h6v-2H9v2Z"/>`),
		g.Group(children),
	)
}