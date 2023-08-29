package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropSixteenNineOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17q-.825 0-1.413-.588T3 15V9q0-.825.588-1.413T5 7h14q.825 0 1.413.588T21 9v6q0 .825-.588 1.413T19 17H5Zm0-2h14V9H5v6Zm0 0V9v6Z"/>`),
		g.Group(children),
	)
}