package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenshotTabletOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20q-.825 0-1.413-.588T1 18V6q0-.825.588-1.413T3 4h18q.825 0 1.413.588T23 6v12q0 .825-.588 1.413T21 20H3ZM4 6H3v12h1V6Zm2 12h12V6H6v12ZM20 6v12h1V6h-1Zm0 0h1h-1ZM4 6H3h1Zm11.5 9.5h-1.75q-.325 0-.537.213T13 16.25q0 .325.213.537t.537.213h2.5q.325 0 .537-.213T17 16.25v-2.5q0-.325-.213-.537T16.25 13q-.325 0-.537.213t-.213.537v1.75Zm-7-7h1.75q.325 0 .537-.213T11 7.75q0-.325-.213-.537T10.25 7h-2.5q-.325 0-.537.213T7 7.75v2.5q0 .325.213.537T7.75 11q.325 0 .537-.213t.213-.537V8.5Z"/>`),
		g.Group(children),
	)
}