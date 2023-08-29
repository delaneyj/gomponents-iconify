package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartHistogramTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" fill-rule="evenodd" d="M4 42H44H4Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M4 42H44"/><rect width="6" height="14" x="8" y="28" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"/><rect width="6" height="24" x="21" y="18" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"/><rect width="6" height="36" x="34" y="6" fill="#2F88FF" stroke="#000" stroke-linejoin="round" stroke-width="4"/></g>`),
		g.Group(children),
	)
}