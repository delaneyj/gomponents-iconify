package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BranchFork(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1946 1472l-365 365l-90-90l210-211H0v-128h384V512h933l-210-211l90-90l365 365l-365 365l-90-90l210-211H512v768h1189l-210-211l90-90l365 365z"/>`),
		g.Group(children),
	)
}