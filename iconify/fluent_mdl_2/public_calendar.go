package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PublicCalendar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1729 1283l317 317l-317 317l-91-90l163-163h-658l163 163l-91 90l-317-317l317-317l91 90l-163 163h658l-163-163l91-90zM256 640v896h512v128H128V128h256V0h128v128h896V0h128v128h256v1024h-128V640H256zm128-384H256v256h1408V256h-128v128h-128V256H512v128H384V256z"/>`),
		g.Group(children),
	)
}