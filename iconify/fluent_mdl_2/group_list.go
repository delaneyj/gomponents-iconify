package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GroupList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M640 640h1408v128H640V640zm1408-384v128H640V256h1408zM384 128v128H256v512h128v128H128V128h256zm256 1408h1408v128H640v-128zm0-384h1408v128H640v-128zm-256-128v128H256v512h128v128H128v-768h256z"/>`),
		g.Group(children),
	)
}