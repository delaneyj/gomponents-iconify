package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayLessonOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22V2h16v9.1q-.45-.05-1-.05t-1 .05V4h-5v7L9.5 9.5L7 11V4H5v16h6.3q.15.5.413 1.038t.537.962H3Zm15 1q-2.075 0-3.538-1.463T13 18q0-2.075 1.463-3.538T18 13q2.075 0 3.538 1.463T23 18q0 2.075-1.463 3.538T18 23Zm-1.25-2.5l4-2.5l-4-2.5v5ZM7 4h5h-5ZM5 4h12h-6h.3H5Z"/>`),
		g.Group(children),
	)
}