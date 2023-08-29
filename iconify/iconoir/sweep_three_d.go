package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SweepThreeD(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 22l5.5-5.5M20 5l-2.5 2.5M7 21h8.5l-7-12l-5.833 10M14.5 3l7 12m-13-6l6-6m1 18l6-6"/>`),
		g.Group(children),
	)
}