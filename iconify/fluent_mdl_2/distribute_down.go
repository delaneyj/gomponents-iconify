package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DistributeDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1747 1587l90 90l-365 365l-365-365l90-90l211 211v-774H512v774l211-211l90 90l-365 365l-365-365l90-90l211 211V896h512V0h128v896h512v902l211-211z"/>`),
		g.Group(children),
	)
}