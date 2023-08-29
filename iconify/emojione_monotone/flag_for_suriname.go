package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForSuriname(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2M9.551 48.717a28.048 28.048 0 0 1-3.375-5.9h51.648a28.023 28.023 0 0 1-3.375 5.9H9.551m44.897-33.434a28.001 28.001 0 0 1 3.376 5.9H6.176a28.001 28.001 0 0 1 3.376-5.9h44.896m-16.37 26.55L32 37.219l-6.076 4.614l2.295-7.506l-6.052-4.65l7.494-.022L32 22.166l2.338 7.488l7.495.022l-6.054 4.65l2.299 7.507"/>`),
		g.Group(children),
	)
}