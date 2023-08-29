package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VibrationOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 15V9h2v6H0Zm3 2V7h2v10H3Zm19-2V9h2v6h-2Zm-3 2V7h2v10h-2ZM8 21q-.825 0-1.413-.588T6 19V5q0-.825.588-1.413T8 3h8q.825 0 1.413.588T18 5v14q0 .825-.588 1.413T16 21H8Zm0-2h8V5H8v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}