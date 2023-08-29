package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForCentralAfricanRepublic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m23.148 14.267A27.834 27.834 0 0 1 60 32H39.867V16.267h15.281M17.531 9.664l.62-1.984l.139-.08l.646 2.064l2.247.006l-1.815 1.396l.688 2.251l-1.823-1.385l-1.825 1.385l.689-2.251l-1.814-1.396l2.248-.006m-8.679 6.603h15.281V32H4a27.834 27.834 0 0 1 4.852-15.733m.024 31.466h15.256v11.134c-6.284-1.844-11.644-5.846-15.256-11.134m30.991 11.134V47.733h15.256c-3.611 5.288-8.971 9.29-15.256 11.134"/>`),
		g.Group(children),
	)
}