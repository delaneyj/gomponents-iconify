package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevenThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M18.02 16.011c0 .735-.398 1.377-.991 1.721v6.275a1 1 0 1 1-2 0v-6.278a1.988 1.988 0 0 1-.736-2.685l-2.143-3.712a1 1 0 1 1 1.732-1l2.13 3.69h.02c1.098 0 1.988.89 1.988 1.99Z"/><path d="M30 16c0 7.732-6.268 14-14 14S2 23.732 2 16S8.268 2 16 2s14 6.268 14 14Zm-3 0c0-6.075-4.925-11-11-11S5 9.925 5 16s4.925 11 11 11s11-4.925 11-11Z"/></g>`),
		g.Group(children),
	)
}