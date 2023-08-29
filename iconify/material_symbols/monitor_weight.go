package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MonitorWeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 12q1.25 0 2.125-.875T15 9q0-1.25-.875-2.125T12 6q-1.25 0-2.125.875T9 9q0 1.25.875 2.125T12 12Zm-1.5-2.5q-.2 0-.35-.15T10 9q0-.2.15-.35t.35-.15q.2 0 .35.15T11 9q0 .2-.15.35t-.35.15Zm1.5 0q-.2 0-.35-.15T11.5 9q0-.2.15-.35T12 8.5q.2 0 .35.15t.15.35q0 .2-.15.35T12 9.5Zm1.5 0q-.2 0-.35-.15T13 9q0-.2.15-.35t.35-.15q.2 0 .35.15T14 9q0 .2-.15.35t-.35.15ZM5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}