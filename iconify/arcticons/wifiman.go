package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifiman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="14.953" cy="31.01" r="2.01" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.19 32.535c5.34-1.347 11.384-2.135 10.953-8.05c4.861 4.474 8.42-.812 12.357-3.277m-12.358 3.276c-5.603-3.4-9.165-.477-11.607 2.06c.095 1.307 3.285.099 6.272-.281"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.473 19.25a3.427 3.427 0 1 1-.055-.112M6.716 23.463a10.59 10.59 0 0 1 14.802-1.272M9.99 26.621a6.148 6.148 0 0 1 8.113-1.2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3.5 19.955c6.212-5.72 13.556-7.022 21.154-1.18"/>`),
		g.Group(children),
	)
}