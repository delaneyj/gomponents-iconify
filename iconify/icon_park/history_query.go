package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HistoryQuery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M42 24V9C42 7.34315 40.6569 6 39 6H9C7.34315 6 6 7.34315 6 9V39C6 40.6569 7.34315 42 9 42H24"/><circle cx="32" cy="32" r="6" fill="#2F88FF"/><path stroke-linecap="round" stroke-linejoin="round" d="M37 36L42 40"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 16H34"/><path stroke-linecap="round" stroke-linejoin="round" d="M14 24L22 24"/></g>`),
		g.Group(children),
	)
}