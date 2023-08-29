package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StayCurrentLandscape(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19q-.825 0-1.413-.588T1 17V7q0-.825.588-1.413T3 5h18q.825 0 1.413.588T23 7v10q0 .825-.588 1.413T21 19H3Zm3-2h12V7H6v10Z"/>`),
		g.Group(children),
	)
}