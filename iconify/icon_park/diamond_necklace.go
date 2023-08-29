package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondNecklace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M17 32.4091L24 29L31 32.4091V39.9091L24 44L17 39.9091V32.4091Z"/><path d="M8 4C8.45455 12.3333 14 29 24 29C34 29 40 12.7838 40 4"/></g>`),
		g.Group(children),
	)
}