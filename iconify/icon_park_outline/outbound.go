package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Outbound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M40 24H24m10-6l6 6l-6 6"/><circle cx="20" cy="24" r="4"/><path stroke-linecap="round" d="M40.706 13A20.102 20.102 0 0 0 38 9.717A19.935 19.935 0 0 0 24 4C12.954 4 4 12.954 4 24s8.954 20 20 20c5.45 0 10.392-2.18 14-5.717A20.104 20.104 0 0 0 40.706 35"/></g>`),
		g.Group(children),
	)
}