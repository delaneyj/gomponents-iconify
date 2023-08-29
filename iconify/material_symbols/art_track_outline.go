package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArtTrackOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19q-.825 0-1.413-.588T1 17V7q0-.825.588-1.413T3 5h10q.825 0 1.413.588T15 7v10q0 .825-.588 1.413T13 19H3Zm0-2h10V7H3v10Zm1-2h8l-2.6-3.5L7.5 14l-1.4-1.85L4 15Zm13 4V5h2v14h-2Zm4 0V5h2v14h-2ZM3 17V7v10Z"/>`),
		g.Group(children),
	)
}