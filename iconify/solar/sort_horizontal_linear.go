package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SortHorizontalLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 8H6m0 0l4.125-4M6 8l4.125 4M6 16h12m0 0l-4.125-4M18 16l-4.125 4"/>`),
		g.Group(children),
	)
}