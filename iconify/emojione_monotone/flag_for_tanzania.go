package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForTanzania(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M32 2C15.432 2 2 15.432 2 32s13.432 30 30 30s30-13.432 30-30S48.568 2 32 2m16.072 7.093c1.148.809 2.232 1.7 3.245 2.667L4.798 38.62a27.932 27.932 0 0 1-.683-4.148L48.072 9.093m11.13 16.287c.328 1.348.558 2.733.683 4.149L15.929 54.907a28.248 28.248 0 0 1-3.246-2.667l46.519-26.86"/>`),
		g.Group(children),
	)
}