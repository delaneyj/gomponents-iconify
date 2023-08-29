package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomToFit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m749 456l-90-91L1024 0l365 365l-90 91l-275-275l-275 275zm550 1136l90 91l-365 365l-365-365l90-91l275 275l275-275zM456 749l-275 275l275 275l-91 90L0 1024l365-365l91 90zm1592 275l-365 365l-91-90l275-275l-275-275l91-90l365 365zM640 640h768v768H640V640zm128 640h512V768H768v512z"/>`),
		g.Group(children),
	)
}