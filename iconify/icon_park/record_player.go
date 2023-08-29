package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordPlayer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><rect width="38" height="32" x="5" y="8" stroke="#000" stroke-width="4" rx="2"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M13 8V40"/><circle cx="28" cy="24" r="9" fill="#2F88FF" stroke="#000" stroke-width="4"/><circle cx="28" cy="24" r="3" fill="#fff"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 16H13"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 24H13"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M5 32H13"/></g>`),
		g.Group(children),
	)
}