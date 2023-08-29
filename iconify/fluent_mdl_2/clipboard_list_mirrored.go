package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardListMirrored(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 256v1792h1536V256h-512q0-53-20-99t-55-82t-81-55t-100-20q-53 0-99 20t-82 55t-55 81t-20 100H256zm1152 256H640V384h256V256q0-27 10-50t27-40t41-28t50-10q27 0 50 10t40 27t28 41t10 50v128h256v128zM384 384h128v256h1024V384h128v1536H384V384zm896 512H512v128h768V896zm0 384H512v128h768v-128zm0 384H512v128h768v-128zm256-768h-128v128h128V896zm0 384h-128v128h128v-128zm0 384h-128v128h128v-128z"/>`),
		g.Group(children),
	)
}