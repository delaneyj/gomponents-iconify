package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1664 256v640q0 116-.5 170.5T1662 1210t-4.5 132t-9.5 108.5t-15.5 102.5t-22.5 84t-31.5 83t-42.5 72H0q79-242 103.5-447.5T128 768V256l256 2v574q0 86 53 139t139 53q139 0 180-93q12-28 12-99V384H640v448q0 39-12 51.5T576 896q-33-3-49-17.5T512 832V256h1152zM527 159q30-50 82-50q60 0 90 24.5t35 82.5h128q0-56-13-89q-20-49-80.5-88T617 0Q517 0 452 55q-68 59-68 147v54h128v-24l1-27l4.5-21l9.5-25z"/>`),
		g.Group(children),
	)
}