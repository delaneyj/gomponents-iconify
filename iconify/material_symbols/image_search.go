package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h5v2H5v14h14v-5.35l2 2V19q0 .825-.588 1.413T19 21H5Zm1-4l3-4l2.25 3l3-4L18 17H6Zm15.55-3.6l-3.1-3.1q-.525.35-1.125.525T16.05 11q-1.85 0-3.15-1.312T11.6 6.5q0-1.875 1.313-3.188T16.1 2q1.875 0 3.188 1.313T20.6 6.5q0 .675-.2 1.3t-.5 1.15L22.95 12l-1.4 1.4ZM16.1 9q1.05 0 1.775-.725T18.6 6.5q0-1.05-.725-1.775T16.1 4q-1.05 0-1.775.725T13.6 6.5q0 1.05.725 1.775T16.1 9Z"/>`),
		g.Group(children),
	)
}