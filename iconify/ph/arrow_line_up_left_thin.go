package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowLineUpLeftThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M220 216a4 4 0 0 1-4 4H40a4 4 0 0 1 0-8h176a4 4 0 0 1 4 4ZM64 156a4 4 0 0 0 4-4V65.66l105.17 105.17a4 4 0 0 0 5.66-5.66L73.66 60H160a4 4 0 0 0 0-8H64a4 4 0 0 0-4 4v96a4 4 0 0 0 4 4Z"/>`),
		g.Group(children),
	)
}