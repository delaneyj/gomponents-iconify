package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaddingRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 0v1920h128V0H0zm2048 0h-128v128h128V0zm0 1792h-128v128h128v-128zm0-1536h-128v256h128V256zm0 384h-128v256h128V640zm0 384h-128v256h128v-256zm0 384h-128v256h128v-256zm-603-512H384v128h1061l-210 211l90 90l365-365l-365-365l-90 90l210 211z"/>`),
		g.Group(children),
	)
}