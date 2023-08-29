package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HowToRegOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-2.8q0-.825.425-1.55t1.175-1.1q1.275-.65 2.875-1.1T10 13q.75 0 1.463.075t1.387.225l-1.75 1.75q-.275-.05-.537-.05H10q-1.775 0-3.188.425T4.5 16.35q-.225.125-.363.35T4 17.2v.8h6.25l2 2H2Zm13.55.4l-3.45-3.45l1.4-1.4l2.05 2.05l5.05-5.05l1.4 1.4l-6.45 6.45ZM10 12q-1.65 0-2.825-1.175T6 8q0-1.65 1.175-2.825T10 4q1.65 0 2.825 1.175T14 8q0 1.65-1.175 2.825T10 12Zm.25 6ZM10 10q.825 0 1.413-.588T12 8q0-.825-.588-1.413T10 6q-.825 0-1.413.588T8 8q0 .825.588 1.413T10 10Zm0-2Z"/>`),
		g.Group(children),
	)
}