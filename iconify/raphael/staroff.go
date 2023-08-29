package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Staroff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 22.375L7.116 28.83l3.396-10.438l-8.883-6.458l10.978.002L16.002 1.5l3.39 10.434h10.982l-8.886 6.457l3.396 10.44L16 22.375zm6.98 3.834l-2.665-8.206l6.98-5.062h-8.628L16 4.73l-2.666 8.205H4.708l6.98 5.07l-2.667 8.203L16 21.146l6.98 5.063z"/>`),
		g.Group(children),
	)
}