package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastRewindRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.975 16.975l-6.2-4.15q-.45-.3-.45-.825t.45-.825l6.2-4.15q.5-.325 1.025-.037t.525.887v8.25q0 .6-.525.9t-1.025-.05Zm-10 0l-6.2-4.15q-.45-.3-.45-.825t.45-.825l6.2-4.15q.5-.325 1.025-.037t.525.887v8.25q0 .6-.525.9t-1.025-.05Z"/>`),
		g.Group(children),
	)
}