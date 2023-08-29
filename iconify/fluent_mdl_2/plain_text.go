package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlainText(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 1535v-255h128v256H128v-1zm384-895v256H384V640h128zm512 1024v-128h128v128h-128zm-128-384h128v256H896v-256zM768 896H640V640h128v256zM640 384v256H512V384h128zm128 512h128v384H256V896h128v256h384V896zM0 1664v-128h128v128H0zm1408-896h384v128h-384V768zm-128 256V896h128v128h-128zm0 256h128v256h-128v-256zm128 384v-128h256v128h-256zm384-768h128v768h-128v-128h-128v-128h128v-128h-384v-128h384V896z"/>`),
		g.Group(children),
	)
}