package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 11V2.1q3.575.375 6.038 2.85T21.9 11H13Zm-1.975 10.875q-3.8-.375-6.363-3.2T2.1 12q0-3.875 2.563-6.7t6.362-3.2v19.775Zm1.975 0v-8.9h8.9q-.35 3.575-2.838 6.063T13 21.875Z"/>`),
		g.Group(children),
	)
}