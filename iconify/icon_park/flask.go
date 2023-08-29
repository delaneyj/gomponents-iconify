package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M37.845 35.0219L44 41.3158V44H4V41.3158L10.1184 35.0594L37.845 35.0219Z"/><path stroke-linecap="round" d="M10.104 35.0743L18.0002 27V6H30.0002V27L37.8727 35.0502"/><path stroke-linecap="round" d="M11 35H37"/><path stroke-linecap="round" d="M30 14H24"/><path stroke-linecap="round" d="M30 20H24"/></g>`),
		g.Group(children),
	)
}