package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalDistributeCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 512v128h-512v256H512V640H0V512h512V256h1024v256h512zm-640-128H640v384h768V384zm384 1024h256v128h-256v256H256v-256H0v-128h256v-256h1536v256zm-128-128H384v384h1280v-384z"/>`),
		g.Group(children),
	)
}