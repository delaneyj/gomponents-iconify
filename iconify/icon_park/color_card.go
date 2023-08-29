package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" d="M10 44C13.3137 44 16 41.3137 16 38V23.5147V4H4V38C4 41.3137 6.68629 44 10 44Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 44C13.3137 44 16 41.3137 16 38V23.5147M10 44C6.68629 44 4 41.3137 4 38V4H16V23.5147M10 44H44V32H24.4853M5.75736 42.2426C8.10051 44.5858 11.8995 44.5858 14.2426 42.2426L24.4853 32M16 23.5147L35.0147 4.5L35.4853 4L43.9853 12.5L24.4853 32"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M14.2427 42.2426L43.9853 12.5L35.4853 4L16 23.5147"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24.4853 32H44V44H12.5"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24.4853 32H44V44H12.5"/><path fill="#2F88FF" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 44C13.3137 44 16 41.3137 16 38V23.5147V4H4V38C4 41.3137 6.68629 44 10 44Z"/></g>`),
		g.Group(children),
	)
}