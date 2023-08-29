package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FastRewindOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.95 16.975l-6.2-4.15q-.45-.3-.45-.825t.45-.825l6.2-4.15q.5-.35 1.025-.05t.525.9v8.25q0 .6-.525.9t-1.025-.05Zm-10 0l-6.2-4.15q-.45-.3-.45-.825t.45-.825l6.2-4.15q.5-.35 1.025-.05t.525.9v8.25q0 .6-.525.9t-1.025-.05ZM9.5 12Zm10 0Zm-10 2.25v-4.5L6.1 12l3.4 2.25Zm10 0v-4.5L16.1 12l3.4 2.25Z"/>`),
		g.Group(children),
	)
}