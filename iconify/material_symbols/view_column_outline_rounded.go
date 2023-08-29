package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumnOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17V7q0-.825.588-1.413T5 5h13.975q.825 0 1.413.588T20.975 7v10q0 .825-.588 1.413T18.976 19H5q-.825 0-1.413-.588T3 17Zm2 0h3.325V7H5v10Zm5.325 0h3.325V7h-3.325v10Zm5.325 0h3.325V7H15.65v10Z"/>`),
		g.Group(children),
	)
}