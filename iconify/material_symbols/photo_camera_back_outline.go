package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCameraBackOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17h12l-3.75-5l-3 4L9 13l-3 4Zm-2 4q-.825 0-1.413-.588T2 19V7q0-.825.588-1.413T4 5h3.15L9 3h6l1.85 2H20q.825 0 1.413.588T22 7v12q0 .825-.588 1.413T20 21H4Zm0-2h16V7h-4.05l-1.825-2h-4.25L8.05 7H4v12Zm8-6Z"/>`),
		g.Group(children),
	)
}