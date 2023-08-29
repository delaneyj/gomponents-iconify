package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartHistogram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M6 6V42H42"/><path d="M14 30V34"/><path d="M22 22V34"/><path d="M30 6V34"/><path d="M38 14V34"/></g>`),
		g.Group(children),
	)
}