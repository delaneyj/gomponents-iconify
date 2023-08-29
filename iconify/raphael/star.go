package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Star(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 22.375L7.116 28.83l3.396-10.438l-8.883-6.458l10.978.002L16.002 1.5l3.39 10.434h10.982l-8.886 6.457l3.396 10.44L16 22.375z"/>`),
		g.Group(children),
	)
}