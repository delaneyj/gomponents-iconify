package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeadphonesRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-7q0-1.875.713-3.513t1.924-2.85q1.213-1.212 2.85-1.924T12 3q1.875 0 3.513.713t2.85 1.924q1.212 1.213 1.925 2.85T21 12v7q0 .825-.588 1.413T19 21h-2q-.825 0-1.413-.588T15 19v-4q0-.825.588-1.413T17 13h2v-1q0-2.925-2.038-4.963T12 5Q9.075 5 7.037 7.038T5 12v1h2q.825 0 1.413.588T9 15v4q0 .825-.588 1.413T7 21H5Z"/>`),
		g.Group(children),
	)
}