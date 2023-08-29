package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ExchangeTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M9 18v24h30V18L24 6L9 18Zm15 12v6m7-10v10m-14-4v4"/><path d="m17 25l5-4l3 3l6-5"/></g>`),
		g.Group(children),
	)
}