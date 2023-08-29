package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneMergeType(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.59 19L7 20.41l6-6V8h3.5L12 3.5L7.5 8H11v5.59zm11.407 1.41l-3.408-3.407l1.4-1.407l3.41 3.408z"/>`),
		g.Group(children),
	)
}