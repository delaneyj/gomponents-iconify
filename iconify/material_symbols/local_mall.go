package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocalMall(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22q-.825 0-1.413-.588T3 20V8q0-.825.588-1.413T5 6h2q0-2.075 1.463-3.538T12 1q2.075 0 3.538 1.463T17 6h2q.825 0 1.413.588T21 8v12q0 .825-.588 1.413T19 22H5Zm7-8q2.075 0 3.538-1.463T17 9h-2q0 1.25-.875 2.125T12 12q-1.25 0-2.125-.875T9 9H7q0 2.075 1.463 3.538T12 14ZM9 6h6q0-1.25-.875-2.125T12 3q-1.25 0-2.125.875T9 6Z"/>`),
		g.Group(children),
	)
}