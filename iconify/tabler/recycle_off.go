package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecycleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 17l-2 2l2 2m-2-2h9m1.896-2.071a2 2 0 0 0-.146-.679l-.55-1M8.536 11l-.732-2.732L5.072 9m2.732-.732l-4.5 7.794a2 2 0 0 0 1.506 2.89l1.141.024M15.464 11l2.732.732L18.928 9m-.732 2.732l-4.5-7.794a2 2 0 0 0-3.256-.14l-.591.976M3 3l18 18"/>`),
		g.Group(children),
	)
}