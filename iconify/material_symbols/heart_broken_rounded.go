package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartBrokenRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.725 19.85Q8 17.45 6.3 15.812t-2.662-2.874q-.963-1.238-1.3-2.263T2 8.5q0-2.3 1.6-3.9T7.5 3q1.125 0 2.175.412T11.55 4.6l-1.175 4.125q-.125.5.163.888t.787.387H13l-.65 6.35q-.025.2.163.225t.237-.15L14.6 10.3q.15-.5-.15-.9t-.8-.4H12l1.775-5.3q.625-.35 1.313-.525T16.5 3q2.3 0 3.9 1.6T22 8.5q0 1.1-.4 2.175t-1.388 2.375q-.987 1.3-2.65 2.938T13.35 19.85q-.275.25-.625.375t-.7.125q-.35 0-.687-.125t-.613-.375Z"/>`),
		g.Group(children),
	)
}