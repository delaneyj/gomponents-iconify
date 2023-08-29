package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FirstQuarterMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M22.5 14a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Z"/><path d="M1 16C1 7.716 7.716 1 16 1c8.284 0 15 6.716 15 15c0 8.284-6.716 15-15 15c-8.284 0-15-6.716-15-15ZM16 4.085a1.5 1.5 0 0 1 0 2.83V19.05a2.505 2.505 0 0 1 1 0a2.5 2.5 0 0 1-1 4.9V29c3.169 0 6.073-1.134 8.328-3.018a3.5 3.5 0 0 1 3.661-4.948A12.96 12.96 0 0 0 29 16c0-7.18-5.82-13-13-13v1.085Zm0 0a1.5 1.5 0 1 0 0 2.83v-2.83Zm0 14.965a2.5 2.5 0 0 0 0 4.9v-4.9Zm-5.283 8.832A3.5 3.5 0 0 0 5.48 23.64a13.05 13.05 0 0 0 5.236 4.24ZM13 12.5a3.5 3.5 0 1 0-7 0a3.5 3.5 0 0 0 7 0Z"/></g>`),
		g.Group(children),
	)
}