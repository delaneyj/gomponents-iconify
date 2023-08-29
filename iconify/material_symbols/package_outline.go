package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PackageOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10 9.75l2-1l2 1V5h-4v4.75ZM7 17v-2h5v2H7Zm-2 4q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5ZM5 5v14V5Zm0 14h14V5h-3v8l-4-2l-4 2V5H5v14Z"/>`),
		g.Group(children),
	)
}