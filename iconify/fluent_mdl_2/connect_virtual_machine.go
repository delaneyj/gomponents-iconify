package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ConnectVirtualMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1792 1280h-768v128h128v128h256v128h384v128h-384v128H512v-128H128v-128h384v-128h256v-128h128v-128H128V128h1664v1152zm-512 384h-256v-128H896v128H640v128h640v-128zM256 1152h1408V256H256v896z"/>`),
		g.Group(children),
	)
}