package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M0 0h128v2048H0V0zm1408 896H256V384h1152v512zm-128-384H384v256h896V512zm640 640v512H256v-512h1664zm-128 128H384v256h1408v-256z"/>`),
		g.Group(children),
	)
}