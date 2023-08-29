package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppBlockingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 23q-.825 0-1.413-.588T5 21V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v4h-2V6H7v12h10v-1h2v4q0 .825-.588 1.413T17 23H7Zm0-2h10v-1H7v1Zm11-5q-1.65 0-2.825-1.175T14 12q0-1.65 1.175-2.825T18 8q1.65 0 2.825 1.175T22 12q0 1.65-1.175 2.825T18 16Zm0-1.5q.3 0 .588-.075t.562-.225l-3.35-3.35q-.15.275-.225.563T15.5 12q0 1.05.725 1.775T18 14.5Zm2.2-1.35q.15-.275.225-.563T20.5 12q0-1.05-.725-1.775T18 9.5q-.3 0-.588.075t-.562.225l3.35 3.35ZM7 4h10V3H7v1Zm0 0V3v1Zm0 17v-1v1Z"/>`),
		g.Group(children),
	)
}