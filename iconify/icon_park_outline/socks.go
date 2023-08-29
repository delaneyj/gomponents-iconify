package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Socks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 18c-3 0-8 0-10 5s.5 9.5 3 12M18 10h22"/><path d="M20 4h18a2 2 0 0 1 2 2v18.288c0 3.432-1.6 6.667-4.336 8.74C32.022 35.788 27.088 39.508 25 41c-3.5 2.5-10 5-15 0c-4.999-5-3.749-11.557.001-15c3.75-3.443 8-6.848 8-6.848V6a2 2 0 0 1 2-2Z"/></g>`),
		g.Group(children),
	)
}