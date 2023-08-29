package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M.546 24.001A.546.546 0 0 1 0 23.455v-1.091c0-.301.245-.545.546-.545h22.909c.301 0 .545.244.546.545v1.091a.546.546 0 0 1-.546.546zm.063-4.364a.675.675 0 0 1-.499-1.052l-.002.002L11.498.291a.574.574 0 0 1 1-.003l.001.003l11.39 18.297a.675.675 0 0 1-.497 1.049h-.003z"/>`),
		g.Group(children),
	)
}