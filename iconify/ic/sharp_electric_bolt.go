package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpElectricBolt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 2L2.5 13L13 14l-5 7l1 1l12.5-11L11 10l5-7z"/>`),
		g.Group(children),
	)
}