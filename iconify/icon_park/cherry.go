package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cherry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><circle cx="14" cy="34" r="8" fill="#2F88FF" stroke-linejoin="round"/><circle cx="35" cy="37" r="7" fill="#2F88FF" stroke-linejoin="round"/><path d="M37 10C34.3488 10.8116 28.6279 13.0145 25.2791 16.2609C20.2558 21.1304 19 24.5 18 27"/><path d="M36.9997 9.99995C35.8831 11.3184 33.7149 14.5959 33.0435 18.3891C32.0364 24.0789 32.9998 27.5 33.9997 29.9999"/><path d="M30 4L44 16"/></g>`),
		g.Group(children),
	)
}