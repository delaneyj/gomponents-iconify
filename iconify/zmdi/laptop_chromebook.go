package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaptopChromebook(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M469 320h43v43H0v-43h43V0h426v320zm-170 0v-21h-86v21h86zm128-64V43H85v213h342z"/>`),
		g.Group(children),
	)
}