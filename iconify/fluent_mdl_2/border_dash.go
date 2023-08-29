package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderDash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 128H0V0h256v128zm384 0H384V0h256v128zm384 0H768V0h256v128zm384 0h-256V0h256v128zm384 0h-256V0h256v128zM0 256h128v256H0V256zm0 384h128v256H0V640zm0 384h128v256H0v-256zm0 384h128v256H0v-256zm0 384h128v256H0v-256zm1792 128h256v128h-256v-128zm-384 0h256v128h-256v-128zm-384 0h256v128h-256v-128zm-384 0h256v128H640v-128zm-384 0h256v128H256v-128zm1664-384h128v256h-128v-256zm0-384h128v256h-128v-256zm0-384h128v256h-128V768zm0-384h128v256h-128V384zM2048 0v256h-128V0h128z"/>`),
		g.Group(children),
	)
}