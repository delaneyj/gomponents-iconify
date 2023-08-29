package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectorOfflineOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.4 21L8 19.6l2.6-2.6L8 14.4L9.4 13l2.6 2.6l2.6-2.6l1.4 1.4l-2.6 2.6l2.6 2.6l-1.4 1.4l-2.6-2.6L9.4 21ZM5 5v1h14V5H5Zm3.1 3l.3 1h7.2l.3-1H8.1Zm.3 3q-.65 0-1.175-.388T6.5 9.6L6 8H5q-.825 0-1.413-.587T3 6V3h18v3q0 .825-.588 1.413T19 8h-1l-.65 1.7q-.225.575-.725.938T15.5 11H8.4ZM5 5v1v-1Z"/>`),
		g.Group(children),
	)
}