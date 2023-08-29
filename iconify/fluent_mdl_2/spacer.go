package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Spacer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m659 1261l90-90l147 146v-293h128v293l147-146l90 90l-301 301l-301-301zm237-365V603L749 749l-90-90l301-301l301 301l-90 90l-147-146v293H896zm1024-768v128H0V128h1920zM0 1664h1920v128H0v-128z"/>`),
		g.Group(children),
	)
}