package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BacklogList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M128 128h1792v512H128V128zm1664 384V256H256v256h1536zM512 1280V768h1408v512H512zm128-384v256h1152V896H640zM128 1920v-512h1792v512H128zm128-384v256h1536v-256H256z"/>`),
		g.Group(children),
	)
}