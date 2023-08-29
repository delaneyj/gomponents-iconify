package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraFrontSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 16h8v-.55q0-1.475-1.413-1.963T10 13q-1.175 0-2.588.488T6 15.45V16Zm4-4q.825 0 1.413-.588T12 10q0-.825-.588-1.413T10 8q-.825 0-1.413.588T8 10q0 .825.588 1.413T10 12Zm-8 8V4h16v6.5l4-4v11l-4-4V20H2Z"/>`),
		g.Group(children),
	)
}