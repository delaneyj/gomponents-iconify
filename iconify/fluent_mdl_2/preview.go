package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Preview(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1115 0l549 549v91h-640V0h91zm37 512h293l-293-293v293zm384 384V768h128v128h-128zm0 256v-128h128v128h-128zm0 256v-128h128v128h-128zm0 256v-128h128v128h-128zm0 256v-128h128v128h-128zm-256 0v-128h128v128h-128zm-256 0v-128h128v128h-128zm-256 0v-128h128v128H768zm-256 0v-128h128v128H512zm-256 0v-128h128v128H256zm0-1024V768h128v128H256zm0-256V512h128v128H256zm0-256V256h128v128H256zm0-256V0h128v128H256zm256 0V0h128v128H512zm256 0V0h128v128H768zM256 1152v-128h128v128H256zm0 256v-128h128v128H256zm0 256v-128h128v128H256z"/>`),
		g.Group(children),
	)
}