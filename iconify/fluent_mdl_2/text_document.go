package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextDocument(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1243 0l549 549v1499H128V0h1115zm37 219v293h293l-293-293zM256 1920h1408V640h-512V128H256v1792zm256-896V896h896v128H512zm0 256v-128h896v128H512zm0 256v-128h896v128H512z"/>`),
		g.Group(children),
	)
}