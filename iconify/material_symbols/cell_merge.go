package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellMerge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-6h2v4h4v2H3Zm12 0v-2h4v-4h2v6h-6Zm-7.825-5.175l-1.425-1.4L7.175 13H2v-2h5.175L5.75 9.575l1.425-1.4L11 12l-3.825 3.825Zm9.65 0L13 12l3.825-3.825l1.425 1.4L16.825 11H22v2h-5.175l1.425 1.425l-1.425 1.4ZM3 9V3h6v2H5v4H3Zm16 0V5h-4V3h6v6h-2Z"/>`),
		g.Group(children),
	)
}