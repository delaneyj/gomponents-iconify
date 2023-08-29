package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraFrontOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 16h8v-.55q0-1.1-1.1-1.775T10 13q-1.8 0-2.9.675T6 15.45V16Zm4-4q.825 0 1.413-.588T12 10q0-.825-.588-1.413T10 8q-.825 0-1.413.588T8 10q0 .825.588 1.413T10 12Zm-6 8q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h12q.825 0 1.413.588T18 6v4.5l4-4v11l-4-4V18q0 .825-.588 1.413T16 20H4Zm0-2h12V6H4v12Zm0 0V6v12Z"/>`),
		g.Group(children),
	)
}