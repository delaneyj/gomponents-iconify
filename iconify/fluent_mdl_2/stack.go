package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 640h256v1152H512v-256H256v-256H0V128h1536v256h256v256zM128 256v896h1280V256H128zm256 1024v128h1280V512h-128v768H384zm1536 384V768h-128v768H640v128h1280z"/>`),
		g.Group(children),
	)
}