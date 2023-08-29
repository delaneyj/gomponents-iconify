package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWakeOnApproach(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 12q-.625 0-1.063-.438T19 10.5v-6q0-.625.438-1.063T20.5 3q.625 0 1.063.438T22 4.5v6q0 .625-.438 1.063T20.5 12ZM10 12q-1.65 0-2.825-1.175T6 8q0-1.65 1.175-2.825T10 4q1.65 0 2.825 1.175T14 8q0 1.65-1.175 2.825T10 12Zm-8 8v-2.8q0-.85.425-1.563T3.6 14.55q1.5-.75 3.112-1.15T10 13q1.675 0 3.288.4t3.112 1.15q.75.375 1.175 1.088T18 17.2V20H2Z"/>`),
		g.Group(children),
	)
}