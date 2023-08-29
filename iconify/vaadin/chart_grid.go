package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartGrid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 9v7h16V9H0zm5 6H1v-1h4v1zm0-2H1v-1h4v1zm0-2H1v-1h4v1zm5 4H6v-1h4v1zm0-2H6v-1h4v1zm0-2H6v-1h4v1zm5 4h-4v-1h4v1zm0-2h-4v-1h4v1zm0-2h-4v-1h4v1zm1-3H0V0h1v7h15v1z"/><path fill="currentColor" d="M15 1.57L9.98 4.43L6.02 2.45L2 4.06v1.08l3.98-1.59l4.04 2.02L15 2.72V1.57z"/>`),
		g.Group(children),
	)
}