package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCameraFrontOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L8.4 3.65q.275-.3.663-.475T9.874 3h4.25q.425 0 .813.175t.662.475L16.85 5H20q.825 0 1.413.588T22 7v12q0 .825-.588 1.413T20 21H4Zm0-2h16V7h-4.05l-1.825-2h-4.25L8.05 7H4v12Zm8-6Zm-4 4h8v-.55q0-1.125-1.1-1.788T12 14q-1.8 0-2.9.663T8 16.45V17Zm4-4q.825 0 1.413-.588T14 11q0-.825-.588-1.413T12 9q-.825 0-1.413.588T10 11q0 .825.588 1.413T12 13Z"/>`),
		g.Group(children),
	)
}