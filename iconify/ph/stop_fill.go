package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 55.27v145.46A15.29 15.29 0 0 1 200.73 216H55.27A15.29 15.29 0 0 1 40 200.73V55.27A15.29 15.29 0 0 1 55.27 40h145.46A15.29 15.29 0 0 1 216 55.27Z"/>`),
		g.Group(children),
	)
}