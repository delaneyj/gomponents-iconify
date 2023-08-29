package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopScreenShare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 23.3L18.2 21H1v-2h15.175l-1-1H4q-.825 0-1.413-.588T2 16V4.85L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM10 12.85L8.175 11q-.125.225-.15.475T8 12v2h2v-1.15Zm10.7 5l-6.275-6.275L16 10l-3-3v2h-1.15l-6-6H20q.825 0 1.413.588T22 5v11q0 .65-.363 1.137t-.937.713Z"/>`),
		g.Group(children),
	)
}