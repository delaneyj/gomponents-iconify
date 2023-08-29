package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BodyFatSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 16h-1v6.2L-.5 12.025L20 1.8V8h1v2h-4V8h1V5.075L13.6 7.25q.675 1.075 1.038 2.275T15 12q0 1.275-.363 2.5t-1.062 2.3l4.4 2.175V16H17v-2h4v2Zm-9.25-.125q.6-.85.925-1.838T13 12q0-1.05-.325-2.013t-.9-1.837L4 12l7.75 3.875Z"/>`),
		g.Group(children),
	)
}