package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EightOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M16.031 18a1.989 1.989 0 0 0 .998-3.71V8.016a1 1 0 1 0-2 0v6.277c-.588.344-.984.98-.987 1.71l-3.68 2.126a1 1 0 1 0 1 1.732l3.679-2.125c.291.168.63.264.99.264Z"/><path d="M30 16c0-7.732-6.268-14-14-14S2 8.268 2 16s6.268 14 14 14s14-6.268 14-14Zm-3 0c0 6.075-4.925 11-11 11S5 22.075 5 16S9.925 5 16 5s11 4.925 11 11Z"/></g>`),
		g.Group(children),
	)
}