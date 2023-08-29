package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElevenOclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M13.882 10.332a1 1 0 1 0-1.732 1l2.143 3.712a1.989 1.989 0 1 0 2.736-.753V8.015a1 1 0 1 0-2 0v4.302l-1.147-1.986Z"/><path d="M30 16c0 7.732-6.268 14-14 14S2 23.732 2 16S8.268 2 16 2s14 6.268 14 14Zm-3 0c0-6.075-4.925-11-11-11S5 9.925 5 16s4.925 11 11 11s11-4.925 11-11Z"/></g>`),
		g.Group(children),
	)
}