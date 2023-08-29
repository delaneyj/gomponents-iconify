package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FontsArt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.57 40.134C8.358 29.57 12.686 5.907 16.727 5.504s5.83 36.996 5.83 36.996"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.57 23.684c4.674-2.135 18.93-6.464 24.356-6.464m3.487 7.162c-3.198-4.796-7.008-2.199-7.18 2.65c-.174 4.848 4.682 4.668 6.471 1.263"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.024 21.837c-1.154 4.329-3.348 15.238.52 15.988s5.887-7.994 5.887-7.994M40.215 13.12a1.219 1.219 0 0 0-1.958-1.293a1.217 1.217 0 0 0-2.402.342c.027 1.226 1.327 2.983 1.327 2.983s2.082-.663 2.821-1.642c.092-.113.165-.243.212-.39Z"/>`),
		g.Group(children),
	)
}