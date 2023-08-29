package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckListTextMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 384h1408v128H0V384zm0 896v-128h1408v128H0zm0 384v-128h1408v128H0zm0-768V768h1408v128H0z"/>`),
		g.Group(children),
	)
}