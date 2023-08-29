package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PolylineOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 22v-2.5L8 16H3v-6h4.3L10 6.9V2h6v6h-4.3L9 11.1v3.15l6 3V16h6v6h-6ZM12 6h2V4h-2v2Zm-7 8h2v-2H5v2Zm12 6h2v-2h-2v2ZM13 5Zm-7 8Zm12 6Z"/>`),
		g.Group(children),
	)
}