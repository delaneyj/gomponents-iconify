package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourThirty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="m17.03 17.731l-.001.001v6.275a1 1 0 1 1-2 0v-6.278a1.988 1.988 0 0 1 1.002-3.707c1.093 0 1.98.881 1.989 1.972l3.679 2.124a1 1 0 1 1-1 1.732l-3.669-2.12Z"/><path d="M30 16c0-7.732-6.268-14-14-14S2 8.268 2 16s6.268 14 14 14s14-6.268 14-14Zm-3 0c0 6.075-4.925 11-11 11S5 22.075 5 16S9.925 5 16 5s11 4.925 11 11Z"/></g>`),
		g.Group(children),
	)
}