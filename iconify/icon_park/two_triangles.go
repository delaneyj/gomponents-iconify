package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoTriangles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><path fill="#2F88FF" d="M24.0002 4L41.3207 34H6.67969L24.0002 4Z"/><path fill="#2F88FF" d="M24.0002 44L41.3207 14H6.67969L24.0002 44Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24.0002 4L41.3207 34H6.67969L24.0002 4Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24.0002 44L41.3207 14H6.67969L24.0002 44Z"/></g>`),
		g.Group(children),
	)
}