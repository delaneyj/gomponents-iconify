package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseWithShieldRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 21q-.825 0-1.413-.588T4 19v-9q0-.475.213-.9t.587-.7l6-4.5q.525-.4 1.2-.4t1.2.4l6 4.5q.375.275.588.7T20 10v9q0 .825-.588 1.413T18 21H6Zm2-8.45q0 1.65.95 3.2t2.575 2.1q.125.05.237.063t.238.012q.125 0 .238-.013t.237-.062q1.625-.55 2.575-2.1t.95-3.2v-1.625q0-.425-.225-.788t-.6-.562l-2.5-1.25q-.325-.15-.675-.15t-.675.15l-2.5 1.25q-.375.2-.6.563T8 10.925v1.625Z"/>`),
		g.Group(children),
	)
}