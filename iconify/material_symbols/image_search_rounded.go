package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageSearchRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h4q.425 0 .713.288T10 4q0 .425-.288.713T9 5H5v14h14v-3.35q0-.425.288-.713T20 14.65q.425 0 .713.288t.287.712V19q0 .825-.588 1.413T19 21H5Zm6.25-5l2.6-3.475q.15-.2.4-.2t.4.2L17.4 16.2q.2.25.05.525T17 17H7q-.3 0-.45-.275t.05-.525l2-2.675q.15-.2.4-.2t.4.2L11.25 16Zm4.8-5q-1.85 0-3.15-1.313T11.6 6.5q0-1.875 1.313-3.188T16.1 2q1.875 0 3.188 1.313T20.6 6.5q0 .675-.2 1.3t-.5 1.15l2.35 2.35q.275.275.275.7t-.275.7q-.275.275-.7.275t-.7-.275l-2.4-2.4q-.525.35-1.125.525T16.05 11Zm.05-2q1.05 0 1.775-.725T18.6 6.5q0-1.05-.725-1.775T16.1 4q-1.05 0-1.775.725T13.6 6.5q0 1.05.725 1.775T16.1 9Z"/>`),
		g.Group(children),
	)
}