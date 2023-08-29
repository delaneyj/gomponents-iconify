package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartStock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><rect width="8" height="16" x="6" y="16" fill="#2F88FF"/><path stroke-linecap="round" d="M10 6V16"/><path stroke-linecap="round" d="M10 32V42"/><rect width="8" height="16" x="34" y="16" fill="#2F88FF"/><path stroke-linecap="round" d="M38 6V16"/><path stroke-linecap="round" d="M38 32V42"/><rect width="8" height="16" x="20" y="14" fill="#2F88FF"/><path stroke-linecap="round" d="M24 4V14"/><path stroke-linecap="round" d="M24 30V40"/></g>`),
		g.Group(children),
	)
}