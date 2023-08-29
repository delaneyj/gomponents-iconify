package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VaccineOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 3l4 4m-2-2l-4.5 4.5m-3-3l6 6m-1-1l-.5.5m-2 2l-4 4H6v-4l4-4m2-2l.5-.5m-5 5L9 14m-6 7l3-3M3 3l18 18"/>`),
		g.Group(children),
	)
}