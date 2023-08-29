package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaddingLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 0v1920h-128V0h128zM0 0h128v128H0V0zm0 1792h128v128H0v-128zM0 256h128v256H0V256zm0 384h128v256H0V640zm0 384h128v256H0v-256zm0 384h128v256H0v-256zm603-512h1061v128H603l210 211l-90 90l-365-365l365-365l90 90l-210 211z"/>`),
		g.Group(children),
	)
}