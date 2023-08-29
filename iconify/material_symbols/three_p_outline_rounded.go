package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThreePOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 14h8v-.55q0-1.1-1.1-1.775T12 11q-1.8 0-2.9.675T8 13.45V14Zm4-4q.825 0 1.413-.588T14 8q0-.825-.588-1.413T12 6q-.825 0-1.413.588T10 8q0 .825.588 1.413T12 10Zm-6 8l-2.3 2.3q-.475.475-1.088.213T2 19.575V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v12q0 .825-.588 1.413T20 18H6Zm-2-2h16V4H4v12Zm0 0V4v12Z"/>`),
		g.Group(children),
	)
}