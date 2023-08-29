package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StockMarket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="8" height="14" x="6" y="20" fill="#2F88FF"/><rect width="8" height="26" x="20" y="14" fill="#2F88FF"/><path stroke-linecap="round" d="M24 44V40"/><rect width="8" height="9" x="34" y="12" fill="#2F88FF"/><path stroke-linecap="round" d="M10 20V10"/><path stroke-linecap="round" d="M38 34V21"/><path stroke-linecap="round" d="M38 12V4"/></g>`),
		g.Group(children),
	)
}