package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpecialEvent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1920 128v1792H0V128h384V0h128v128h896V0h128v128h384zM128 256v256h1664V256h-256v128h-128V256H512v128H384V256H128zm1664 1536V640H128v1152h1664zm-440-768l-241 189l101 315l-252-197l-252 197l101-315l-241-189h302l90-280l90 280h302z"/>`),
		g.Group(children),
	)
}