package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="6.617" height="8.768" x="10.5" y="27.213" rx="3.309" ry="3.309"/><path d="M27.304 35.981v-5.46a3.309 3.309 0 0 0-3.308-3.308h0a3.309 3.309 0 0 0-3.309 3.309m0 5.459v-8.768m16.379 7.099a3.308 3.308 0 0 1-2.875 1.67h0a3.309 3.309 0 0 1-3.308-3.31v-2.15a3.309 3.309 0 0 1 3.308-3.309h0a3.309 3.309 0 0 1 3.309 3.309v1.075h-6.617"/><circle cx="29.62" cy="16.959" r="4.94"/></g><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}