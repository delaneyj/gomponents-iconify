package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MemoryCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 40H8C6.89543 40 6 39.1046 6 38V19.1056C6 18.7107 6.11688 18.3247 6.3359 17.9962L12.4063 8.8906C12.7772 8.3342 13.4017 8 14.0704 8H40C41.1046 8 42 8.89543 42 10V38C42 39.1046 41.1046 40 40 40Z"/><path d="M18 16L18 8"/><path d="M24 16L24 8"/><path d="M30 16L30 8"/><path d="M36 16L36 8"/><path fill="#2F88FF" d="M15 28H33V40H15V28Z"/></g>`),
		g.Group(children),
	)
}