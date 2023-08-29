package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventAccepted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M2048 128v1024h-128V640H128v1152h896v128H0V128h384V0h128v128h1024V0h128v128h384zm-128 384V256h-256v128h-128V256H512v128H384V256H128v256h1792zm19 851l90 90l-557 558l-269-270l90-90l179 178l467-466z"/>`),
		g.Group(children),
	)
}