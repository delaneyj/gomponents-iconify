package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareChevUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.35 13.15a.495.495 0 0 1-.7.7L12 11.21l-2.65 2.64a.495.495 0 0 1-.7-.7l3-3a.492.492 0 0 1 .7 0Z"/><path fill="currentColor" d="M20.937 5.563v12.874a2.5 2.5 0 0 1-2.5 2.5H5.563a2.5 2.5 0 0 1-2.5-2.5V5.563a2.5 2.5 0 0 1 2.5-2.5h12.874a2.5 2.5 0 0 1 2.5 2.5ZM4.063 18.437a1.5 1.5 0 0 0 1.5 1.5h12.874a1.5 1.5 0 0 0 1.5-1.5V5.563a1.5 1.5 0 0 0-1.5-1.5H5.563a1.5 1.5 0 0 0-1.5 1.5Z"/>`),
		g.Group(children),
	)
}