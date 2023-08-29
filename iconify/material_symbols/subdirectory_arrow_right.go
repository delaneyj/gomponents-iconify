package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SubdirectoryArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14 21l-1.425-1.425l3.6-3.575H5V4h2v10h9.175l-3.6-3.6l1.4-1.425L20 15l-6 6Z"/>`),
		g.Group(children),
	)
}