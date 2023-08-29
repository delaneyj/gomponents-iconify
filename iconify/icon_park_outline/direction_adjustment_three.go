package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DirectionAdjustmentThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="m18 10l6-6m0 0l6 6m-6-6v20m0 0v20m0-20h20m-20 0H4m14 14l6 6m0 0l6-6m8-20l6 6m0 0l-6 6M10 18l-6 6m0 0l6 6"/>`),
		g.Group(children),
	)
}