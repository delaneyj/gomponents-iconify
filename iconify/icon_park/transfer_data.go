package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransferData(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M10 8L4 14L10 20"/><path d="M38 28L44 34L38 40"/><path d="M4 14H44"/><path d="M4 34H44"/></g>`),
		g.Group(children),
	)
}