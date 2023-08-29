package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlaceItem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V9q0-.825.588-1.413T5 7h4v2H5v10h14V9h-4V7h4q.825 0 1.413.588T21 9v10q0 .825-.588 1.413T19 21H5Zm7-5l-4-4l1.4-1.4l1.6 1.575V0h2v12.175l1.6-1.575L16 12l-4 4Z"/>`),
		g.Group(children),
	)
}