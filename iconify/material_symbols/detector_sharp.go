package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DetectorSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21q-2.2 0-4.213-.825t-3.562-2.4L5.65 16.35q1.275 1.275 2.925 1.963t3.45.687q1.8 0 3.438-.675t2.912-1.95l1.4 1.4q-1.575 1.575-3.575 2.4T12 21Zm0-4q-1.4 0-2.675-.525T7.05 14.95l1.4-1.4q.725.725 1.638 1.088T12 15q1 0 1.913-.363t1.637-1.087l1.4 1.4q-1 1-2.275 1.525T12 17ZM8.1 8l.3 1h7.2l.3-1H8.1Zm-1.15 3L6 8H3V3h18v5h-3l-1.15 3h-9.9Z"/>`),
		g.Group(children),
	)
}