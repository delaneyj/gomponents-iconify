package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EntitlementRedemption(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 512h640v128H256V512zm1152 256v128H256V768h1152zm0 256v128H256v-128h1152zm0 256v128H256v-128h1152zM256 1664v-128h926l-128 128H256zm-128 256h1027l128 128H0V0h1115l549 549v862l-128 128V640h-512V128H128v1792zM1152 219v293h293l-293-293zm787 1144l90 90l-557 558l-269-270l90-90l179 178l467-466z"/>`),
		g.Group(children),
	)
}