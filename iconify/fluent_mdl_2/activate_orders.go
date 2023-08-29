package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ActivateOrders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2029 1453l-557 558l-269-270l90-90l179 178l467-466l90 90zM1024 640H640V512h384v128zm0 256H640V768h384v128zm-384 128h384v128H640v-128zM512 640H384V512h128v128zm0 256H384V768h128v128zm-128 128h128v128H384v-128zm768-384V128H256v1792h896v128H128V0h1115l549 549v731l-128 128V640h-512zm128-128h293l-293-293v293z"/>`),
		g.Group(children),
	)
}