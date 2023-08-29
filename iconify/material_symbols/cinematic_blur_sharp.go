package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CinematicBlurSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4 3l2 4h3L7 3h2l2 4h3l-2-4h2l2 4h3l-2-4h5v18H2V3h2Zm4 15h8v-.55q0-1.1-1.1-1.775T12 15q-1.8 0-2.9.675T8 17.45V18Zm4-4q.825 0 1.413-.588T14 12q0-.825-.588-1.413T12 10q-.825 0-1.413.588T10 12q0 .825.588 1.413T12 14Z"/>`),
		g.Group(children),
	)
}