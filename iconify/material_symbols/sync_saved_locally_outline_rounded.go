package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncSavedLocallyOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.925 11.225L9.5 9.8q-.15-.15-.325-.225T8.812 9.5q-.187 0-.374.075T8.1 9.8q-.275.275-.275.7t.275.7l2.125 2.15q.3.3.7.3t.7-.3l4.65-4.25q0-.4-.05-.775t-.325-.65q-.3-.3-.713-.3t-.712.3l-3.55 3.55ZM1 21v-2h21q.425 0 .713.288T23 20q0 .425-.288.713T22 21H1Zm3-3q-.825 0-1.413-.588T2 16V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v11q0 .825-.588 1.413T20 18H4Zm0-2h16V5H4v11Zm0 0V5v11Z"/>`),
		g.Group(children),
	)
}