package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CropRotateSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 24q-2.35 0-4.438-.838t-3.7-2.324q-1.612-1.488-2.637-3.5T0 13h2q.2 1.8.963 3.363t1.974 2.762q1.213 1.2 2.788 1.938t3.375.887L9.55 20.4l1.4-1.4l4.5 4.5q-.85.275-1.712.388T12 24Zm3-5v-2H7V9H5V7h2V5h2v10h10v2h-2v2h-2Zm0-6V9h-4V7h6v6h-2Zm7-2q-.175-1.8-.95-3.363t-1.988-2.762q-1.212-1.2-2.787-1.938T12.9 2.05l1.55 1.55l-1.4 1.4L8.55.5Q9.4.225 10.263.112T12 0q2.35 0 4.438.838t3.7 2.325q1.612 1.487 2.637 3.5T24 11h-2Z"/>`),
		g.Group(children),
	)
}