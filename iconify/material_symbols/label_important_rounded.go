package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabelImportantRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 12L4.975 8.075q-.65-1-.075-2.037T6.675 5H15q.5 0 .938.225t.712.625l3.525 5q.375.525.375 1.15t-.375 1.15l-3.525 5q-.275.4-.712.625T15 19H6.675q-1.2 0-1.775-1.038t.075-2.037L7.5 12Z"/>`),
		g.Group(children),
	)
}