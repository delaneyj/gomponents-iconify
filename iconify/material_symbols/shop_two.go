package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 15l5.5-3.5L12 8v7Zm-9 7q-.825 0-1.413-.588T1 20V9h2v11h17v2H3Zm4-4q-.825 0-1.413-.588T5 16V5h5V3q0-.825.588-1.413T12 1h4q.825 0 1.413.588T18 3v2h5v11q0 .825-.588 1.413T21 18H7Zm5-13h4V3h-4v2Z"/>`),
		g.Group(children),
	)
}