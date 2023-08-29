package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LabPanelOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19v-5q0-.825.588-1.413T3 12V7.725q-.45-.275-.725-.712T2 6V5q0-.825.588-1.413T4 3h5q.825 0 1.413.588T11 5v1q0 .575-.275 1.012T10 7.726V12h4V7.725q-.45-.275-.725-.712T13 6V5q0-.825.588-1.413T15 3h5q.825 0 1.413.588T22 5v1q0 .575-.275 1.012T21 7.726V12q.825 0 1.413.588T23 14v5q0 .825-.588 1.413T21 21H3ZM15 6h5V5h-5v1ZM4 6h5V5H4v1Zm12 6h3V8h-3v4ZM5 12h3V8H5v4Zm-2 7h18v-5H3v5ZM4 6V5v1Zm11 0V5v1ZM3 19v-5v5Z"/>`),
		g.Group(children),
	)
}