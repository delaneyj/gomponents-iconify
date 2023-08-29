package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BodyFat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 16h-1v2.975q0 .575-.263 1t-.687.7q-.425.275-.938.3T17.1 20.75l-14-6.95q-.575-.275-.838-.763T2 12.026Q2 11.5 2.263 11t.837-.775l14-7q.5-.25 1.012-.213t.938.313q.425.275.688.7t.262 1V8h1v2h-4V8h1V5.075L13.6 7.25q.675 1.075 1.038 2.275T15 12q0 1.275-.363 2.5t-1.062 2.3l4.4 2.175V16H17v-2h4v2Zm-9.25-.125q.6-.85.925-1.838T13 12q0-1.05-.325-2.013t-.9-1.837L4 12l7.75 3.875Z"/>`),
		g.Group(children),
	)
}