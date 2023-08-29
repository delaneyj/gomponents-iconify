package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExpandCategories(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 26h6v2h-6zm0-8h8v2h-8zm0-8h10v2H20zm-5-6h2v24h-2zm-4.414-.041L7 7.249L3.412 3.958L2 5.373L7 10l5-4.627l-1.414-1.414z"/>`),
		g.Group(children),
	)
}