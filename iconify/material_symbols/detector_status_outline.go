package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectorStatusOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.95 21L7.4 17.45l1.425-1.4l2.125 2.125l4.225-4.25L16.6 15.35L10.95 21ZM5 5v1h14V5H5Zm3.1 3l.3 1h7.2l.3-1H8.1Zm.3 3q-.65 0-1.175-.388T6.5 9.6L6 8H5q-.825 0-1.413-.587T3 6V3h18v3q0 .825-.588 1.413T19 8h-1l-.65 1.7q-.225.575-.725.938T15.5 11H8.4ZM5 5v1v-1Z"/>`),
		g.Group(children),
	)
}