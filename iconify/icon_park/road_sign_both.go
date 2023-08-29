package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RoadSignBoth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M10 8V16H38L42 12L38 8H10Z"/><path fill="#2F88FF" d="M38 23V31H10L6 27L10 23H38Z"/><path stroke-linecap="round" d="M24 31V44"/><path stroke-linecap="round" d="M24 16V23"/><path stroke-linecap="round" d="M24 4V8"/><path stroke-linecap="round" d="M19 44H29"/></g>`),
		g.Group(children),
	)
}