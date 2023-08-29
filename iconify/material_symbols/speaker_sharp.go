package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 22V2h14v20H5Zm7-13q.825 0 1.413-.588T14 7q0-.825-.588-1.413T12 5q-.825 0-1.413.588T10 7q0 .825.588 1.413T12 9Zm0 10q1.65 0 2.825-1.175T16 15q0-1.65-1.175-2.825T12 11q-1.65 0-2.825 1.175T8 15q0 1.65 1.175 2.825T12 19Zm0-2q-.825 0-1.413-.588T10 15q0-.825.588-1.413T12 13q.825 0 1.413.588T14 15q0 .825-.588 1.413T12 17Z"/>`),
		g.Group(children),
	)
}