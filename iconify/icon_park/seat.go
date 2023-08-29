package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Seat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" d="M17 21.458C12.9429 22.7323 10 26.5226 10 31.0002C10 36.5231 14.4772 41.0002 20 41.0002C23.2716 41.0002 26.1763 39.4291 28.0007 37.0002C28.2404 36.6811 28.4615 36.3471 28.6623 36"/><path stroke-linecap="round" stroke-linejoin="round" d="M38 20C34 20 30.5 19.5 24 17V29H38V43"/><circle cx="24" cy="8" r="4" fill="#2F88FF"/></g>`),
		g.Group(children),
	)
}