package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M17.029 8.016a1 1 0 1 0-2 0v6.277A1.988 1.988 0 0 0 16.03 18c1.099 0 1.989-.89 1.989-1.989V16l3.648-2.106a1 1 0 0 0-1-1.732l-3.64 2.1V8.017Z"/><path d="M30 16c0-7.732-6.268-14-14-14S2 8.268 2 16s6.268 14 14 14s14-6.268 14-14Zm-3 0c0 6.075-4.925 11-11 11S5 22.075 5 16S9.925 5 16 5s11 4.925 11 11Z"/></g>`),
		g.Group(children),
	)
}