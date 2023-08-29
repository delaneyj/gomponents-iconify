package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareChevDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.65 10.85a.495.495 0 0 1 .7-.7L12 12.79l2.65-2.64a.495.495 0 0 1 .7.7l-3 3a.492.492 0 0 1-.7 0Z"/><path fill="currentColor" d="M3.063 18.437V5.563a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5v12.874a2.5 2.5 0 0 1-2.5 2.5H5.563a2.5 2.5 0 0 1-2.5-2.5ZM19.937 5.563a1.5 1.5 0 0 0-1.5-1.5H5.563a1.5 1.5 0 0 0-1.5 1.5v12.874a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5Z"/>`),
		g.Group(children),
	)
}