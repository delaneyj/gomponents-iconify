package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadsetMicRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 23q-.425 0-.713-.288T12 22q0-.425.288-.713T13 21h6v-1h-2q-.825 0-1.413-.588T15 18v-4q0-.825.588-1.413T17 12h2v-1q0-2.9-2.05-4.95T12 4Q9.1 4 7.05 6.05T5 11v1h2q.825 0 1.413.588T9 14v4q0 .825-.588 1.413T7 20H5q-.825 0-1.413-.588T3 18v-7q0-1.85.713-3.488T5.65 4.65q1.225-1.225 2.863-1.938T12 2q1.85 0 3.488.713T18.35 4.65q1.225 1.225 1.938 2.863T21 11v10q0 .825-.588 1.413T19 23h-6Z"/>`),
		g.Group(children),
	)
}