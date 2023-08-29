package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompareRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 21H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h5V2q0-.425.288-.713T11 1q.425 0 .713.288T12 2v20q0 .425-.288.713T11 23q-.425 0-.713-.288T10 22v-1Zm-5-3h5v-6l-5 6Zm9 3v-9l5 6V5h-5V3h5q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21h-5Z"/>`),
		g.Group(children),
	)
}