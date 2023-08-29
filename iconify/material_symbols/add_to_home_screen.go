package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AddToHomeScreen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 23q-.825 0-1.413-.588T6 21v-4h2v1h10V6H8v1H6V3q0-.825.588-1.413T8 1h10q.825 0 1.413.588T20 3v18q0 .825-.588 1.413T18 23H8Zm-3.6-6L3 15.6L8.6 10H5V8h7v7h-2v-3.6L4.4 17Z"/>`),
		g.Group(children),
	)
}