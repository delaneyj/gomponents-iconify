package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCameraFrontOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 21V5h5.15L9 3h6l1.85 2H22v16H2Zm2-2h16V7h-4.05l-1.825-2h-4.25L8.05 7H4v12Zm8-6Zm-4 4h8v-.55q0-1.125-1.1-1.788T12 14q-1.8 0-2.9.663T8 16.45V17Zm4-4q.825 0 1.413-.588T14 11q0-.825-.588-1.413T12 9q-.825 0-1.413.588T10 11q0 .825.588 1.413T12 13Z"/>`),
		g.Group(children),
	)
}