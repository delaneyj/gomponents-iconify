package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OutpatientMedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.5 15.5l-1.425-1.4l1.1-1.1H16v-2h3.15l-1.075-1.075L19.5 8.5L23 12l-3.5 3.5ZM2 5V3h12v2H2Zm4.5 12.5h3V15H12v-3H9.5V9.5h-3V12H4v3h2.5v2.5ZM3 21q-.825 0-1.413-.588T1 19V8q0-.825.588-1.413T3 6h10q.825 0 1.413.588T15 8v11q0 .825-.588 1.413T13 21H3Zm0-2h10V8H3v11Zm5-5.5Z"/>`),
		g.Group(children),
	)
}