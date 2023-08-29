package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NineThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M18.02 16.011c0 .735-.398 1.377-.991 1.721v6.275a1 1 0 1 1-2 0v-6.278a2 2 0 0 1-.724-.729h-4.272a1 1 0 0 1 0-2h4.285a1.988 1.988 0 0 1 3.702 1.011Z"/><path d="M30 16c0-7.732-6.268-14-14-14S2 8.268 2 16s6.268 14 14 14s14-6.268 14-14Zm-3 0c0 6.075-4.925 11-11 11S5 22.075 5 16S9.925 5 16 5s11 4.925 11 11Z"/></g>`),
		g.Group(children),
	)
}