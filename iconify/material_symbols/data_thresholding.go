package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataThresholding(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm10.15-2h2.125L19 17.275V16h-.85l-3 3Zm-7.5-5l3.025-3l2 2l5.075-5.1l-1.4-1.4l-3.675 3.675l-2-2L6.25 12.6l1.4 1.4ZM5 19h.85l3-3H6.725L5 17.725V19Zm8.525 0l3-3H14.4l-3 3h2.125ZM9.8 19l3-3h-2.125l-3 3H9.8Z"/>`),
		g.Group(children),
	)
}