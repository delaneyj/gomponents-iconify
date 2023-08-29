package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MedicationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 17.5h3V15H16v-3h-2.5V9.5h-3V12H8v3h2.5v2.5ZM7 21q-.825 0-1.413-.588T5 19V8q0-.825.588-1.413T7 6h10q.825 0 1.413.588T19 8v11q0 .825-.588 1.413T17 21H7Zm0-2h10V8H7v11ZM6 5V3h12v2H6Zm1 3v11V8Z"/>`),
		g.Group(children),
	)
}