package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PageCheckedOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M256 128v1792h1152v128H128V0h1115l549 549v923l-128-128V640h-512V128H256zm1024 91v293h293l-293-293zm640 1317h128v512h-512v-128h293l-402-403l90-90l403 402v-293z"/>`),
		g.Group(children),
	)
}