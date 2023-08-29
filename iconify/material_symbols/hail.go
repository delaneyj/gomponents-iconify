package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22v-6h3v6H4Zm5 0V10.1q-1.25.425-1.625 1.563T7 14H5q0-3.2 1.875-5.1T12 7q2.5 0 3.75-1.238T17 2h2q0 2.2-.938 3.938T15 8.4V22h-2v-6h-2v6H9Zm3-16q-.825 0-1.413-.588T10 4q0-.825.588-1.413T12 2q.825 0 1.413.588T14 4q0 .825-.588 1.413T12 6Z"/>`),
		g.Group(children),
	)
}