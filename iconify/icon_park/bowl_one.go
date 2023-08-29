package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BowlOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke-linejoin="round" d="M24 39C33.3888 39 41 31.9411 41 22H7C7 31.9411 14.6112 39 24 39Z"/><path stroke-linejoin="round" d="M18 38L16 44H32L30 38"/><path d="M12 10L12 14"/><path d="M36 10L36 14"/><path d="M24 5L24 14"/></g>`),
		g.Group(children),
	)
}