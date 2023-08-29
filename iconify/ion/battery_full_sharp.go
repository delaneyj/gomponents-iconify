package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryFullSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M17 384h432V128H17Zm32-224h368v192H49Z"/><path fill="currentColor" d="M70.69 182.94h324.63v146.13H70.69zM465 202.67h32v106.67h-32z"/>`),
		g.Group(children),
	)
}