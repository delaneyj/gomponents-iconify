package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1664h-768v384H896v-384H128v-512h768V896H384V384h512V0h128v384h512v512h-512v256h768v512zM512 768h896V512H512v256zm1152 512H256v256h1408v-256z"/>`),
		g.Group(children),
	)
}