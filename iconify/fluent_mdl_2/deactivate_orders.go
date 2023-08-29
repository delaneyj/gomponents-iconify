package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeactivateOrders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1819 1728l226 227l-90 90l-227-226l-227 227l-90-91l227-227l-227-227l90-90l227 227l227-227l90 91l-226 226zM512 640H384V512h128v128zm0 256H384V768h128v128zm-128 128h128v128H384v-128zm-128 896h1024v128H128V0h1115l549 549v731h-128V640h-512V128H256v1792zM1280 512h293l-293-293v293zm-256 128H640V512h384v128zm0 256H640V768h384v128zm-384 128h384v128H640v-128z"/>`),
		g.Group(children),
	)
}