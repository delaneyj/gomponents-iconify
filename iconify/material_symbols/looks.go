package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Looks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 17q0-2.275.863-4.275t2.362-3.5q1.5-1.5 3.5-2.363T12 6q2.275 0 4.275.863t3.5 2.362q1.5 1.5 2.363 3.5T23 17h-2q0-3.725-2.638-6.363T12 8q-3.725 0-6.363 2.638T3 17H1Zm4 0q0-2.9 2.05-4.95T12 10q2.9 0 4.95 2.05T19 17h-2q0-2.075-1.463-3.538T12 12q-2.075 0-3.538 1.463T7 17H5Z"/>`),
		g.Group(children),
	)
}