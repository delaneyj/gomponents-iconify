package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LastQuarterMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M13 12.5a3.5 3.5 0 1 1-7 0a3.5 3.5 0 0 1 7 0Z"/><path d="M16 1C7.716 1 1 7.716 1 16c0 8.284 6.716 15 15 15c8.284 0 15-6.716 15-15c0-8.284-6.716-15-15-15ZM3 16C3 8.82 8.82 3 16 3v1.085a1.498 1.498 0 0 0-1 0a1.5 1.5 0 0 0 1 2.83V19.05a2.5 2.5 0 1 1 0 4.9v-4.9a2.5 2.5 0 0 0 0 4.9V29c-1.881 0-3.669-.4-5.283-1.118A3.5 3.5 0 0 0 5.48 23.64A12.942 12.942 0 0 1 3 16Zm21.328 9.982a3.5 3.5 0 0 1 3.661-4.948a13.036 13.036 0 0 1-3.66 4.948ZM16 6.915v-2.83a1.5 1.5 0 0 1 0 2.83Zm9 4.585a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/></g>`),
		g.Group(children),
	)
}