package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalculatorGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1024 1280v-128h128v128h-128zm0 256v-128h128v128h-128zm384 256v-128h128v128h-128zm0-256v-128h128v128h-128zm-384 256v-128h128v128h-128zM1536 0v128H256v1664H128V0h1408zM768 1152v128H640v-128h128zm0 256v128H640v-128h128zm0 256v128H640v-128h128zM384 256h1408v1792H384V256zm1280 1664V384H512v1536h1152zM1536 512v384H640V512h896zm-128 256V640H768v128h640zm0 512v-128h128v128h-128z"/>`),
		g.Group(children),
	)
}