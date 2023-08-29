package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestCamIq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 22q-.825 0-1.413-.588T7 20h4v-4.075q-2.6-.375-4.3-2.337T5 9q0-2.925 2.038-4.963T12 2q2.925 0 4.963 2.038T19 9q0 2.625-1.7 4.588T13 15.925V20h4q0 .825-.588 1.413T15 22H9Zm3-8q2.075 0 3.538-1.463T17 9q0-2.075-1.463-3.538T12 4Q9.925 4 8.462 5.463T7 9q0 2.075 1.463 3.538T12 14Zm0-3q-.825 0-1.413-.588T10 9q0-.825.588-1.413T12 7q.825 0 1.413.588T14 9q0 .825-.588 1.413T12 11Z"/>`),
		g.Group(children),
	)
}