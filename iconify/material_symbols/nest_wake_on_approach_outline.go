package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWakeOnApproachOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 12q-.625 0-1.063-.438T19 10.5v-6q0-.625.438-1.063T20.5 3q.625 0 1.063.438T22 4.5v6q0 .625-.438 1.063T20.5 12ZM10 12q-1.65 0-2.825-1.175T6 8q0-1.65 1.175-2.825T10 4q1.65 0 2.825 1.175T14 8q0 1.65-1.175 2.825T10 12Zm-8 8v-2.8q0-.85.425-1.563T3.6 14.55q1.5-.75 3.112-1.15T10 13q1.675 0 3.288.4t3.112 1.15q.75.375 1.175 1.088T18 17.2V20H2Zm2-2h12v-.8q0-.275-.125-.5t-.375-.35q-1.275-.65-2.663-1T10 15q-1.45 0-2.838.35t-2.662 1q-.25.125-.375.35T4 17.2v.8Zm6-8q.825 0 1.413-.588T12 8q0-.825-.588-1.413T10 6q-.825 0-1.413.588T8 8q0 .825.588 1.413T10 10Zm0-2Zm0 10Z"/>`),
		g.Group(children),
	)
}