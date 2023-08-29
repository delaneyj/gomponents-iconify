package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 256v1536H0V256h2048zm-128 128h-384v256h-128V384H640v256H512V384H128v1280h384v-128h128v128h768v-128h128v128h384V384zM512 768h128v256H512V768zm0 384h128v256H512v-256zm896-384h128v256h-128V768zm0 384h128v256h-128v-256z"/>`),
		g.Group(children),
	)
}